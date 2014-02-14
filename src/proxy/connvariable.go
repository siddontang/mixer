package proxy

import (
	"fmt"
	. "mysql"
	. "parser"
	"strconv"
	"strings"
)

func (c *conn) handleSetVariable(l *lex) error {
	switch strings.ToUpper(l.Get(1).Value) {
	case `@@AUTOCOMMIT`, `AUTOCOMMIT`:
		return c.handleSetAutoCommit(l)
	default:
		return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("set %s can not supported now", l.Get(1).Value))
	}
}

func (c *conn) handleSetAutoCommit(l *lex) error {
	if l.Get(2).Type == TK_EQ &&
		l.Get(3).Type == TK_INTEGER &&
		(l.Get(4).Type == TK_EOF || l.Get(4).Type == TK_SEMICOLON) {
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
		return NewError(ER_UNKNOWN_ERROR, "set autocommit error")
	}

	return c.writeOK(nil)
}
