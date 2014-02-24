package proxy

import (
	"fmt"
	. "mysql"
	. "parser"
	"strconv"
	"strings"
)

func (c *conn) handleSet(l *lex) error {
	switch strings.ToUpper(l.Get(1).Value) {
	case `AUTOCOMMIT`:
		return c.handleSetAutoCommit(l)
	case `NAMES`:
		return c.handleSetNames(l)
	default:
		return fmt.Errorf("set %s can not supported now", l.Get(1).Value)
	}
}

func (c *conn) handleSetAutoCommit(l *lex) error {
	if l.Get(2).Type == TK_EQ &&
		l.Get(3).Type == TK_INTEGER {
		if i, err := strconv.Atoi(l.Get(3).Value); err != nil {
			return err
		} else {
			if i == 0 {
				c.status &= ^SERVER_STATUS_AUTOCOMMIT
			} else if i == 1 {
				c.status |= SERVER_STATUS_AUTOCOMMIT
			} else {
				return NewDefaultError(ER_WRONG_VALUE_FOR_VAR, "autocommit", i)
			}
		}
	} else {
		return fmt.Errorf("set autocommit error")
	}

	return c.writeOK(nil)
}

func (c *conn) handleSetNames(l *lex) error {
	if strings.ToUpper(l.Get(3).Value) == "COLLATE" {
		return NewError(ER_UNKNOWN_ERROR, "set collate not supported now")
	}

	charset := strings.Trim(l.Get(2).Value, "\"'`")
	cid, ok := CharsetIds[charset]
	if !ok {
		return fmt.Errorf("invalid charset %s", charset)
	}

	c.charset = charset
	c.collation = cid

	return c.writeOK(nil)
}
