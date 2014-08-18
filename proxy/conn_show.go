package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/hack"
	. "github.com/siddontang/mixer/mysql"
	"github.com/siddontang/mixer/sqlparser"
	"strings"
)

func (c *Conn) handleShow(sql string, stmt *sqlparser.Show) error {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Name) {
	case "databases":
		r, err = c.handleShowDatabases()
	default:
		err = fmt.Errorf("unsupport show %s now", sql)
	}

	if err != nil {
		return err
	}

	return c.writeResultset(c.status, r)
}

func (c *Conn) handleShowDatabases() (*Resultset, error) {
	dbs := make([]interface{}, 0, len(c.server.schemas))
	for key := range c.server.schemas {
		dbs = append(dbs, key)
	}

	return c.buildShowResult(dbs, "Database")
}

func (c *Conn) buildShowResult(values []interface{}, name string) (*Resultset, error) {
	r := new(Resultset)

	field := &Field{}

	field.Name = hack.Slice(name)
	field.Charset = 33
	field.Type = MYSQL_TYPE_VAR_STRING

	r.Fields = []*Field{field}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}
		r.RowDatas = append(r.RowDatas,
			PutLengthEncodedString(row))
	}

	return r, nil
}
