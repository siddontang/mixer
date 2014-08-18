package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/hack"
	. "github.com/siddontang/mixer/mysql"
	"github.com/siddontang/mixer/sqlparser"
	"strconv"
	"strings"
)

func (c *Conn) handleSimpleSelect(sql string, stmt *sqlparser.SimpleSelect) error {
	if len(stmt.SelectExprs) != 1 {
		return fmt.Errorf("support select one informaction function, %s", sql)
	}

	expr, ok := stmt.SelectExprs[0].(*sqlparser.NonStarExpr)
	if !ok {
		return fmt.Errorf("support select informaction function, %s", sql)
	}

	var f *sqlparser.FuncExpr
	f, ok = expr.Expr.(*sqlparser.FuncExpr)
	if !ok {
		return fmt.Errorf("support select informaction function, %s", sql)
	}

	var r *Resultset
	var err error

	switch strings.ToLower(string(f.Name)) {
	case "last_insert_id":
		r, err = c.buildSimpleSelectResult(c.lastInsertId, f.Name, expr.As)
	case "row_count":
		r, err = c.buildSimpleSelectResult(c.affectedRows, f.Name, expr.As)
	case "version":
		r, err = c.buildSimpleSelectResult(ServerVersion, f.Name, expr.As)
	case "connection_id":
		r, err = c.buildSimpleSelectResult(c.connectionId, f.Name, expr.As)
	case "database":
		if c.schema != nil {
			r, err = c.buildSimpleSelectResult(c.schema.db, f.Name, expr.As)
		} else {
			r, err = c.buildSimpleSelectResult("NULL", f.Name, expr.As)
		}
	default:
		return fmt.Errorf("function %s not support", f.Name)
	}

	if err != nil {
		return err
	}

	return c.writeResultset(c.status, r)
}

func formatValue(value interface{}) ([]byte, error) {
	switch v := value.(type) {
	case int8:
		return strconv.AppendInt(nil, int64(v), 10), nil
	case int16:
		return strconv.AppendInt(nil, int64(v), 10), nil
	case int32:
		return strconv.AppendInt(nil, int64(v), 10), nil
	case int64:
		return strconv.AppendInt(nil, int64(v), 10), nil
	case int:
		return strconv.AppendInt(nil, int64(v), 10), nil
	case uint8:
		return strconv.AppendUint(nil, uint64(v), 10), nil
	case uint16:
		return strconv.AppendUint(nil, uint64(v), 10), nil
	case uint32:
		return strconv.AppendUint(nil, uint64(v), 10), nil
	case uint64:
		return strconv.AppendUint(nil, uint64(v), 10), nil
	case uint:
		return strconv.AppendUint(nil, uint64(v), 10), nil
	case float32:
		return strconv.AppendFloat(nil, float64(v), 'f', -1, 64), nil
	case float64:
		return strconv.AppendFloat(nil, float64(v), 'f', -1, 64), nil
	case []byte:
		return v, nil
	case string:
		return hack.Slice(v), nil
	default:
		return nil, fmt.Errorf("invalid type %T", value)
	}
}

//fields must have name set
func (c *Conn) buildResultset(fields []*Field, values [][]interface{}) (*Resultset, error) {
	r := new(Resultset)

	r.Fields = fields

	var row []byte
	var b []byte
	var err error

	for i, vs := range values {
		if len(vs) != len(fields) {
			return nil, fmt.Errorf("row %d has %d column not equal %d", i, len(vs), len(fields))
		}

		row = row[0:0]
		for j, value := range vs {
			if i == 0 {
				field := fields[j]
				if field.Name == nil {
					return nil, fmt.Errorf("field %d must set name", j)
				}

				switch value.(type) {
				case int8, int16, int32, int64, int:
					field.Charset = 63
					field.Type = MYSQL_TYPE_LONGLONG
					field.Flag = BINARY_FLAG | NOT_NULL_FLAG
				case uint8, uint16, uint32, uint64, uint:
					field.Charset = 63
					field.Type = MYSQL_TYPE_LONGLONG
					field.Flag = BINARY_FLAG | NOT_NULL_FLAG | UNSIGNED_FLAG
				case string, []byte:
					field.Charset = 33
					field.Type = MYSQL_TYPE_VAR_STRING
				default:
					return nil, fmt.Errorf("unsupport type %T for resultset", value)
				}
			} else {
				switch value.(type) {
				case int8, int16, int32, int64, int:
					if r.Fields[j].Type != MYSQL_TYPE_LONGLONG {
						return nil, fmt.Errorf("invalid type %T at (%d, %d), must int", value, i, j)
					}
				case uint8, uint16, uint32, uint64, uint:
					if r.Fields[j].Type != MYSQL_TYPE_LONGLONG {
						return nil, fmt.Errorf("invalid type %T at (%d, %d), must int", value, i, j)
					}
				case string, []byte:
					if r.Fields[j].Type != MYSQL_TYPE_VAR_STRING {
						return nil, fmt.Errorf("invalid type %T at (%d, %d), must string", value, i, j)
					}
				default:
					return nil, fmt.Errorf("unsupport type %T for resultset", value)
				}

			}

			b, err = formatValue(value)

			if err != nil {
				return nil, err
			}

			row = append(row, PutLengthEncodedString(b)...)
		}

		r.RowDatas = append(r.RowDatas, row)
	}

	return r, nil
}

func (c *Conn) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*Resultset, error) {

	field := &Field{}

	field.Name = name

	if asName != nil {
		field.Name = asName
	}

	field.OrgName = name

	return c.buildResultset([]*Field{field}, [][]interface{}{[]interface{}{value}})
}
