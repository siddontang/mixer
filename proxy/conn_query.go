package proxy

import (
	"fmt"
	. "github.com/siddontang/mixer/mysql"
	"github.com/siddontang/mixer/sqlparser"
)

func (c *Conn) handleQuery(sql string) error {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return err
	}

	switch v := stmt.(type) {
	case *sqlparser.Select:
	case *sqlparser.Insert:
	case *sqlparser.Update:
	case *sqlparser.Delete:
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

func (c *Conn) handleExecute(stmt sqlparser.Statement) error {
	if c.schema == nil {
		return NewDefaultError(ER_NO_DB_ERROR)
	}

	nodes, err := sqlparser.GetStmtShardList(stmt, nil, c.schema.rule)
	if err != nil {
		return err
	}

	return nil
}
