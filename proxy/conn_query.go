package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/client"
	. "github.com/siddontang/mixer/mysql"
	"github.com/siddontang/mixer/sqlparser"
)

func (c *Conn) handleQuery(sql string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("execute %s error %v", sql, e)
			return
		}
	}()

	var stmt sqlparser.Statement
	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		return err
	}

	switch v := stmt.(type) {
	case *sqlparser.Select:
		return c.handleSelect(sql, stmt)
	case *sqlparser.Insert:
		return c.handleExec(sql, stmt)
	case *sqlparser.Update:
		return c.handleExec(sql, stmt)
	case *sqlparser.Delete:
		return c.handleExec(sql, stmt)
	case *sqlparser.Set:
		return c.handleSet(v)
	case *sqlparser.Begin:
		return c.handleBegin()
	case *sqlparser.Commit:
		return c.handleCommit()
	case *sqlparser.Rollback:
		return c.handleRollback()
	default:
		return fmt.Errorf("statement %T not support now", stmt)
	}

	return nil
}

func (c *Conn) getShardList(stmt sqlparser.Statement) ([]*Node, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	ns, err := sqlparser.GetStmtShardList(stmt, nil, c.schema.rule)
	if err != nil {
		return nil, err
	}

	if len(ns) == 0 {
		return nil, fmt.Errorf("must shard to a node")
	}

	n := make([]*Node, 0, len(ns))
	for _, name := range ns {
		n = append(n, c.server.getNode(name))
	}
	return n, nil
}

func (c *Conn) getConn(n *Node, isSelect bool) (co *client.SqlConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.getSelectConn()
		} else {
			co, err = n.getMasterConn()
		}
		if err != nil {
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.getMasterConn(); err != nil {
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
	if err = co.UseDB(c.schema.db); err != nil {
		return
	}

	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

func (c *Conn) handleSelect(sql string, stmt sqlparser.Statement) error {
	nodes, err := c.getShardList(stmt)
	if err != nil {
		return err
	}

	conns := make([]*client.SqlConn, 0, len(nodes))

	var co *client.SqlConn
	for _, n := range nodes {
		co, err = c.getConn(n, true)
		if err != nil {
			break
		}

		conns = append(conns, co)
	}

	return nil
}

func (c *Conn) handleExec(sql string, stmt sqlparser.Statement) error {
	if !c.isInTransaction() {

	}

	return nil
}
