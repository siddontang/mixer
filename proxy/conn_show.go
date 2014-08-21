package proxy

import (
	"fmt"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/mixer/hack"
	. "github.com/siddontang/mixer/mysql"
	"github.com/siddontang/mixer/sqlparser"
	"sort"
	"strings"
	"time"
)

func (c *Conn) handleShow(sql string, stmt *sqlparser.Show) error {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Section) {
	case "databases":
		r, err = c.handleShowDatabases()
	case "tables":
		r, err = c.handleShowTables(sql, stmt)
	case "proxy":
		r, err = c.handleShowProxy(sql, stmt)
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

	return c.buildSimpleShowResultset(dbs, "Database")
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

	return c.buildSimpleShowResultset(values, fmt.Sprintf("Tables_in_%s", s.db))
}

func (c *Conn) handleShowProxy(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Key) {
	case "config":
		r, err = c.handleShowProxyConfig()
	case "status":
		r, err = c.handleShowProxyStatus(sql, stmt)
	default:
		err = fmt.Errorf("Unsupport show proxy [%v] yet, just support [config|status] now.", stmt.Key)
		log.Warn(err.Error())
		return nil, err
	}
	return r, err
}

func (c *Conn) handleShowProxyConfig() (*Resultset, error) {
	var names []string = []string{"Section", "Key", "Value"}
	var rows [][]string
	const (
		Column = 3
	)

	rows = append(rows, []string{"Global_Config", "Addr", c.server.cfg.Addr})
	rows = append(rows, []string{"Global_Config", "User", c.server.cfg.User})
	rows = append(rows, []string{"Global_Config", "Password", c.server.cfg.Password})
	rows = append(rows, []string{"Global_Config", "LogLevel", c.server.cfg.LogLevel})
	rows = append(rows, []string{"Global_Config", "Schemas_Count", fmt.Sprintf("%d", len(c.server.schemas))})
	rows = append(rows, []string{"Global_Config", "Nodes_Count", fmt.Sprintf("%d", len(c.server.nodes))})

	for db, schema := range c.server.schemas {
		rows = append(rows, []string{"Schemas", "DB", db})

		var nodeNames []string
		var nodeRows [][]string
		for name, node := range schema.nodes {
			nodeNames = append(nodeNames, name)
			var nodeSection = fmt.Sprintf("Schemas[%s]-Node[ %v ]", db, name)

			if node.master != nil {
				nodeRows = append(nodeRows, []string{nodeSection, "Master", node.master.String()})
			}
			if node.masterBackup != nil {
				nodeRows = append(nodeRows, []string{nodeSection, "Master_Backup", node.masterBackup.String()})
			}

			if node.slave != nil {
				nodeRows = append(nodeRows, []string{nodeSection, "Slave", node.slave.String()})
			}
			nodeRows = append(nodeRows, []string{nodeSection, "Last_Master_Ping", fmt.Sprintf("%v", time.Unix(node.lastMasterPing, 0))})

			nodeRows = append(nodeRows, []string{nodeSection, "Last_Slave_Ping", fmt.Sprintf("%v", time.Unix(node.lastSlavePing, 0))})

			nodeRows = append(nodeRows, []string{nodeSection, "down_after_noalive", fmt.Sprintf("%v", node.downAfterNoAlive)})

		}
		rows = append(rows, []string{fmt.Sprintf("Schemas[%s]", db), "Nodes_List", strings.Join(nodeNames, ",")})

		var defaultRule = schema.rule.DefaultRule
		if defaultRule.DB == db {
			if defaultRule.DB == db {
				rows = append(rows, []string{fmt.Sprintf("Schemas[%s]_Rule_Default", db),
					"Default_Table", defaultRule.String()})
			}
		}
		for tb, r := range schema.rule.Rules {
			if r.DB == db {
				rows = append(rows, []string{fmt.Sprintf("Schemas[%s]_Rule_Table", db),
					fmt.Sprintf("Table[ %s ]", tb), r.String()})
			}
		}

		rows = append(rows, nodeRows...)

	}

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(names, values)
}

func (c *Conn) handleShowProxyStatus(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	// TODO: handle like_or_where expr
	return nil, nil
}

func (c *Conn) buildSimpleShowResultset(values []interface{}, name string) (*Resultset, error) {

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
