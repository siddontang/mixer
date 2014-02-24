package proxy

import (
	"errors"
	"fmt"
	"lib/log"
	. "mysql"
	. "parser"
	"strings"
)

type lex struct {
	Query    string
	Tokens   []Token
	Args     []interface{} //for stmt args
	Prepared bool          //check prepare statement
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

	l := &lex{query, tokens, nil, false}

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
		return c.handleSet(l)
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
		return NewError(ER_UNKNOWN_ERROR,
			fmt.Sprintf("command %s not supported now", l.Get(0).Value))
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

	for _, co := range c.txConns {
		if e := co.Commit(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = map[*node]*Conn{}

	return
}

func (c *conn) rollback() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Rollback(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = map[*node]*Conn{}

	return
}

type routeFunc func(name string, co *Conn, query string, args ...interface{}) interface{}

//if status is in_trans, need
//else if status is not autocommit, need
//else no need
func (c *conn) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}

func (c *conn) getDBConn(n *node) (co *Conn, err error) {
	if !c.needBeginTx() {
		co, err = n.GetConn()
		if err != nil {
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.GetConn(); err != nil {
				return
			}

			if err = co.Begin(); err != nil {
				return
			}

			c.Lock()
			c.txConns[n] = co
			c.Unlock()
		}
	}

	//todo, set conn charset, etc...
	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

func (c *conn) route(l *lex, f routeFunc) ([]interface{}, error) {
	if c.curSchema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	rs, err := c.curSchema.Route(l, c.needBeginTx())
	if err != nil {
		log.Error("schema route error %s", err.Error())
		return nil, err
	}

	ch := make(chan interface{}, len(rs))

	for n, query := range rs {
		go func(n *node, q routeQuery, f routeFunc) {
			if co, err := c.getDBConn(n); err != nil {
				ch <- err
			} else {
				ch <- f(n.Name, co, q.Query, q.Args...)
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

func routeExec(nodeName string, co *Conn, query string, args ...interface{}) interface{} {
	r, err := co.Exec(query, args...)
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

	c.affectedRows = int64(-1)

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
			if r.InsertId == 0 {
				r.InsertId = v.InsertId
			} else if r.InsertId > v.InsertId {
				//last insert id is first gen id for multi row inserted
				//see http://dev.mysql.com/doc/refman/5.6/en/information-functions.html#function_last-insert-id
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
		if r.InsertId > 0 {
			c.lastInsertId = int64(r.InsertId)
		}

		c.affectedRows = int64(r.AffectedRows)

		return c.writeOK(r)
	}
}
