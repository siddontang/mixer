package proxy

import (
	"bytes"
	"errors"
	"fmt"
	"lib/log"
	. "mysql"
	. "parser"
	"strings"
)

type lex struct {
	Query  string
	Tokens []Token
	Args   []interface{} //for stmt args
}

func (l *lex) Get(index int) Token {
	if index >= 0 && index < len(l.Tokens) {
		return l.Tokens[index]
	} else {
		return Token{TK_UNKNOWN, ""}
	}
}

func parseQuery(query string) (*lex, error) {
	tokens, err := Tokenizer(query)

	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, errors.New("No Token")
	}

	l := &lex{query, tokens, nil}

	return l, nil
}

func (c *conn) handleQuery(data []byte) error {
	l, err := parseQuery(string(data))
	if err != nil {
		return err
	}

	switch l.Get(0).Type {
	case TK_SQL_SELECT:
		return c.handleSelect(l)
	case TK_SQL_INSERT:
		return c.handleExec(l)
	case TK_SQL_UPDATE:
		return c.handleExec(l)
	case TK_SQL_DELETE:
		return c.handleExec(l)
	case TK_SQL_REPLACE:
		return c.handleExec(l)
	case TK_SQL_SET:
		return c.handleSetVariable(l)
	default:
		return c.handleQueryLiteral(l)
	}

	return nil
}

func (c *conn) handleQueryLiteral(l *lex) error {
	switch strings.ToUpper(l.Get(0).Value) {
	case `BEGIN`:
		return c.handleBegin()
	case `COMMIT`:
		return c.handleCommit()
	case `ROLLBACK`:
		return c.handleRollback()
	default:
		return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("command %s not supported now", l.Get(0).Value))
	}
}

func (c *conn) isInTransaction() bool {
	return c.status&SERVER_STATUS_IN_TRANS > 0
}

func (c *conn) isAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *conn) handleBegin() error {
	c.status |= SERVER_STATUS_IN_TRANS
	return c.writeOK(nil)
}

func (c *conn) handleCommit() (err error) {
	if err := c.commit(); err != nil {
		return err
	} else {
		return c.writeOK(nil)
	}
}

func (c *conn) handleRollback() (err error) {
	if err := c.rollback(); err != nil {
		return err
	} else {
		return c.writeOK(nil)
	}
}

func (c *conn) commit() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, tx := range c.txs {
		if e := tx.Commit(); e != nil {
			err = e
		}
	}

	c.txs = map[*node]*Tx{}

	return
}

func (c *conn) rollback() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, tx := range c.txs {
		if e := tx.Rollback(); e != nil {
			err = e
		}
	}
	return
}

type dbQueryer interface {
	Prepare(query string) (*Stmt, error)
	Query(query string, args ...interface{}) (*Resultset, error)
	Exec(query string, args ...interface{}) (*Result, error)
}

type routeFunc func(name string, q dbQueryer, query string, args ...interface{}) interface{}

//if status is in_trans, need
//else if status is not autocommit, need
//else no need
func (c *conn) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}

func (c *conn) getDBQueryer(n *node) (dbQueryer, error) {
	if !c.needBeginTx() {
		return n, nil
	} else {
		c.Lock()
		tx, ok := c.txs[n]
		c.Unlock()

		if ok {
			return tx, nil
		}

		var err error
		if tx, err = n.Begin(); err != nil {
			log.Error("node %s begin error %s", n.Name, err.Error())
			return nil, err
		}

		c.Lock()
		c.txs[n] = tx
		c.Unlock()

		return tx, nil
	}
}

func (c *conn) route(l *lex, f routeFunc) ([]interface{}, error) {
	if c.curSchema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	rs, err := c.curSchema.Route(l)
	if err != nil {
		log.Error("schema route error %s", err.Error())
		return nil, NewError(ER_UNKNOWN_ERROR, err.Error())
	}

	ch := make(chan interface{}, len(rs))

	for n, query := range rs {
		go func(n *node, q routeQuery, f routeFunc) {
			if d, err := c.getDBQueryer(n); err != nil {
				ch <- err
			} else {
				ch <- f(n.Name, d, q.Query, q.Args...)
			}
		}(n, query, f)
	}

	results := make([]interface{}, 0, len(rs))
LOOP:
	for {
		select {
		case r := <-ch:
			results = append(results, r)
			if len(results) == len(rs) {
				break LOOP
			}
		}
	}

	return results, nil
}

func routeSelect(nodeName string, q dbQueryer, query string, args ...interface{}) interface{} {
	r, err := q.Query(query, args...)
	if err != nil {
		log.Error("node %s query error %s", nodeName, err.Error())
		return err
	} else {
		return r
	}
}

func mergeResultset(dest *Resultset, src *Resultset) error {
	if dest.ColumnNumber() != src.ColumnNumber() {
		return errors.New("column not match")
	}

	for i := range dest.Fields {
		//here we test name, type and flag
		if !bytes.Equal(dest.Fields[i].Name, src.Fields[i].Name) {
			return fmt.Errorf("field name %s != %s", dest.Fields[i].Name, src.Fields[i].Name)
		}

		if dest.Fields[i].Type != src.Fields[i].Type {
			return fmt.Errorf("field type %d != %d", dest.Fields[i].Type, src.Fields[i].Type)
		}

		if dest.Fields[i].Flag != src.Fields[i].Flag {
			return fmt.Errorf("field flag %d != %d", dest.Fields[i].Flag, src.Fields[i].Flag)
		}
	}

	dest.Status |= src.Status

	//later we may merge with select condition like limit, order by, etc...
	//now we only append row
	for _, v := range src.RowPackets {
		dest.RowPackets = append(dest.RowPackets, v)
	}

	return nil
}

func (c *conn) writeResultset(r *Resultset) error {
	columnLen := PutLengthEncodedInt(uint64(len(r.Fields)))

	data := make([]byte, 4, 1024)

	data = append(data, columnLen...)
	if err := c.WritePacket(data); err != nil {
		return err
	}

	for _, v := range r.Fields {
		data = data[0:4]
		data = append(data, v.Packet...)
		if err := c.WritePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(r.Status); err != nil {
		return err
	}

	for _, v := range r.RowPackets {
		data = data[0:4]
		data = append(data, v...)
		if err := c.WritePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(r.Status); err != nil {
		return err
	}

	return nil
}

func (c *conn) handleSelect(l *lex) error {
	results, err := c.route(l, routeSelect)
	if err != nil {
		return err
	}

	var r *Resultset = nil

LOOP:
	for _, i := range results {
		switch v := i.(type) {
		case error:
			err = v
			break LOOP
		case *Resultset:
			if r == nil {
				r = v
			} else {
				if e := mergeResultset(r, v); err != nil {
					err = e
					break LOOP
				}
			}
		default:
			err = fmt.Errorf("invalid return type %T", i)
			break LOOP
		}
	}

	if err != nil {
		return err
	} else {
		r.Status |= c.status
		return c.writeResultset(r)
	}
}

func routeExec(nodeName string, q dbQueryer, query string, args ...interface{}) interface{} {
	r, err := q.Exec(query, args...)
	if err != nil {
		log.Error("node %s exec error %s", nodeName, err.Error())
		return err
	} else {
		return r
	}
}

func (c *conn) handleExec(l *lex) error {
	results, err := c.route(l, routeExec)
	if err != nil {
		return err
	}

	var r = new(Result)
	r.Status = c.status

LOOP:
	for _, i := range results {
		switch v := i.(type) {
		case error:
			err = v
			break LOOP
		case (*Result):
			r.Status |= v.Status
			r.AffectedRows += v.AffectedRows
			if r.InsertId < v.InsertId {
				r.InsertId = v.InsertId
			}
		default:
			err = fmt.Errorf("invalid return type %T", i)
			break LOOP
		}
	}

	if err != nil {
		return err
	} else {
		return c.writeOK(r)
	}
}
