package proxy

// import (
//     "errors"
//     "fmt"
//     "github.com/siddontang/go-log/log"
//     "github.com/siddontang/mixer/client"
//     . "github.com/siddontang/mixer/mysql"
//     . "github.com/siddontang/mixer/parser"
//     "strconv"
//     "strings"
// )

// func (c *conn) writeResultset(r *Resultset) error {
//     r.Status |= c.status

//     c.affectedRows = int64(-1)

//     columnLen := PutLengthEncodedInt(uint64(len(r.Fields)))

//     data := make([]byte, 4, 1024)

//     data = append(data, columnLen...)
//     if err := c.writePacket(data); err != nil {
//         return err
//     }

//     for _, v := range r.Fields {
//         data = data[0:4]
//         data = append(data, v.Dump()...)
//         if err := c.writePacket(data); err != nil {
//             return err
//         }
//     }

//     if err := c.writeEOF(r.Status); err != nil {
//         return err
//     }

//     for _, v := range r.RowDatas {
//         data = data[0:4]
//         data = append(data, v...)
//         if err := c.writePacket(data); err != nil {
//             return err
//         }
//     }

//     if err := c.writeEOF(r.Status); err != nil {
//         return err
//     }

//     return nil
// }

// func formatString(value interface{}) (string, error) {
//     switch v := value.(type) {
//     case int8:
//         return strconv.FormatInt(int64(v), 10), nil
//     case int16:
//         return strconv.FormatInt(int64(v), 10), nil
//     case int32:
//         return strconv.FormatInt(int64(v), 10), nil
//     case int64:
//         return strconv.FormatInt(int64(v), 10), nil
//     case int:
//         return strconv.FormatInt(int64(v), 10), nil
//     case uint8:
//         return strconv.FormatUint(uint64(v), 10), nil
//     case uint16:
//         return strconv.FormatUint(uint64(v), 10), nil
//     case uint32:
//         return strconv.FormatUint(uint64(v), 10), nil
//     case uint64:
//         return strconv.FormatUint(uint64(v), 10), nil
//     case uint:
//         return strconv.FormatUint(uint64(v), 10), nil
//     case float32:
//         return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
//     case float64:
//         return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
//     case []byte:
//         return string(v), nil
//     case string:
//         return v, nil
//     default:
//         return "", fmt.Errorf("invalid type %T", value)
//     }
// }

// //handle select informaction function, like select last_insert_id();
// func (c *conn) selectInfoFunc(l *lex, value interface{}) (*Resultset, error) {
//     if l.Prepared {
//         return nil, fmt.Errorf("prepare statment not supported now")
//     }

//     if !(l.Get(2).Type == TK_LPAREN && l.Get(3).Type == TK_RPAREN) {
//         return nil, fmt.Errorf("only select function(); supported now")
//     }

//     name := l.Get(1).Value
//     if l.Get(4).Type == TK_SQL_AS {
//         alias := l.Get(5).Value
//         if len(alias) == 0 {
//             return nil, fmt.Errorf("some error around as %s", l.Get(5).Value)
//         } else {
//             name = alias
//         }
//     }

//     r := new(Resultset)

//     field := &Field{}

//     field.Name = []byte(name)
//     field.OrgName = []byte(l.Get(1).Value)

//     var row string

//     switch v := value.(type) {
//     case int8, int16, int32, int64, int:
//         field.Charset = 63
//         field.Type = MYSQL_TYPE_LONGLONG
//         field.Flag = BINARY_FLAG | NOT_NULL_FLAG
//         row, _ = formatString(value)
//     case uint8, uint16, uint32, uint64, uint:
//         field.Charset = 63
//         field.Type = MYSQL_TYPE_LONGLONG
//         field.Flag = BINARY_FLAG | NOT_NULL_FLAG | UNSIGNED_FLAG
//         row, _ = formatString(value)
//     case string:
//         field.Charset = 33
//         field.Type = MYSQL_TYPE_VAR_STRING
//         row = v
//     default:
//         return nil, fmt.Errorf("unsupport type %T for resultset", value)
//     }

//     r.Fields = []*Field{field}

//     r.RowDatas = append(r.RowDatas,
//         PutLengthEncodedString([]byte(row)))

//     return r, nil
// }

// func routeSelect(nodeName string, co *client.SqlConn, query string, args ...interface{}) interface{} {
//     r, err := co.Query(query, args...)
//     if err != nil {
//         log.Error("node %s query error %s", nodeName, err.Error())
//         return err
//     } else {
//         return r
//     }
// }

// func mergeResultset(dest *Resultset, src *Resultset) error {
//     if dest.ColumnNumber() != src.ColumnNumber() {
//         return errors.New("column not match")
//     }

//     dest.Status |= src.Status

//     //later we may merge with select condition like limit, order by, etc...
//     //now we only append row
//     for _, v := range src.RowDatas {
//         dest.RowDatas = append(dest.RowDatas, v)
//     }

//     return nil
// }

// func (c *conn) handleSelectDefault(l *lex) (*Resultset, error) {

//     results, err := c.route(l, routeSelect)
//     if err != nil {
//         return nil, err
//     }

//     var r *Resultset = nil

// LOOP:
//     for _, i := range results {
//         switch v := i.(type) {
//         case error:
//             err = v
//             break LOOP
//         case *Resultset:
//             if r == nil {
//                 r = v
//             } else {
//                 if e := mergeResultset(r, v); err != nil {
//                     err = e
//                     break LOOP
//                 }
//             }
//         default:
//             err = fmt.Errorf("invalid return type %T", i)
//             break LOOP
//         }
//     }

//     return r, err
// }

// func (c *conn) handleSelect(l *lex) error {
//     var r *Resultset
//     var err error
//     switch strings.ToUpper(l.Get(1).Value) {
//     case `LAST_INSERT_ID`:
//         //only support select last_insert_id(), not last_insert_id(expr)
//         r, err = c.selectInfoFunc(l, c.lastInsertId)
//     case `ROW_COUNT`:
//         r, err = c.selectInfoFunc(l, c.affectedRows)
//     case `VERSION`:
//         r, err = c.selectInfoFunc(l, ServerVersion)
//     case `DATABASE`:
//         if c.curSchema != nil {
//             r, err = c.selectInfoFunc(l, c.curSchema.db)
//         } else {
//             r, err = c.selectInfoFunc(l, "NULL")
//         }
//     case `CONNECTION_ID`:
//         r, err = c.selectInfoFunc(l, c.connectionId)
//     default:
//         r, err = c.handleSelectDefault(l)
//     }

//     if err != nil {
//         return err
//     } else {
//         return c.writeResultset(r)
//     }
// }
