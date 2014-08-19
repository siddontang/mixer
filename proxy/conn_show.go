package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/hack"
	. "github.com/siddontang/mixer/mysql"
	"github.com/siddontang/mixer/sqlparser"
	"sort"
	"strings"
)

func (c *Conn) handleShow(sql string, stmt *sqlparser.Show) error {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Section) {
	case "databases":
		r, err = c.handleShowDatabases()
	case "tables":
		r, err = c.handleShowTables(sql, stmt)
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

	return c.buildShowResultset(dbs, "Database")
}

func (c *Conn) handleShowTables(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	s := c.schema
	if stmt.From != nil {
		db := nstring(stmt.From)
		s = c.server.getSchema(db)
	}

	if s == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	var tables []string
	tmap := map[string]struct{}{}
	for _, n := range s.nodes {
		co, err := n.getMasterConn()
		if err != nil {
			return nil, err
		}

		if err := co.UseDB(s.db); err != nil {
			co.Close()
			return nil, err
		}

		if r, err := co.Execute(sql); err != nil {
			co.Close()
			return nil, err
		} else {
			co.Close()
			for i := 0; i < r.RowNumber(); i++ {
				n, _ := r.GetString(i, 0)
				if _, ok := tmap[n]; !ok {
					tables = append(tables, n)
				}
			}
		}
	}

	sort.Strings(tables)

	values := make([]interface{}, len(tables))
	for i := range tables {
		values[i] = tables[i]
	}

	return c.buildShowResultset(values, fmt.Sprintf("Tables_in_%s", s.db))
}

func (c *Conn) buildShowResultset(values []interface{}, name string) (*Resultset, error) {

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
