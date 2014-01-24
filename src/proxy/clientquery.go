package proxy

import (
	"bytes"
	"fmt"
	"github.com/siddontang/golib/log"
	"unicode"
)

func (c *ClientConn) getQueryCmd(data []byte) (string, error) {
	//trim left blank
	buf := bytes.TrimLeftFunc(data, unicode.IsSpace)

	pos := bytes.IndexFunc(buf, unicode.IsSpace)
	if pos == -1 {
		pos = len(buf)
	}

	return string(bytes.TrimRight(buf[0:pos], "; \t\n")), nil
}

func (c *ClientConn) handleQuery(data []byte) error {
	//trim left blank
	cmd, err := c.getQueryCmd(data)
	if err != nil {
		return err
	}

	log.Info("query %s", data)

	switch cmd {
	case "select":
		return c.handleSelect(data)
	case "update":
		return c.handleExec(data)
	case "insert":
		return c.handleExec(data)
	case "delete":
		return c.handleExec(data)
	case "replace":
		return c.handleExec(data)
	case "begin":
		return c.handleBegin()
	case "commit":
		return c.handleCommit()
	case "rollback":
		return c.handleRollback()
	default:
		return NewMySQLError(ER_UNKNOWN_ERROR, fmt.Sprintf("command %s not supported now", data))
	}

	return nil
}

func (c *ClientConn) isInTrans() bool {
	return c.status&SERVER_STATUS_IN_TRANS > 0
}

func (c *ClientConn) routeQuery(data []byte) error {
	if c.schema == nil {
		return NewDefaultMySQLError(ER_NO_DB_ERROR)
	}

	r, err := c.schema.Route(data)
	if err != nil {
		log.Error("schema route error %s", err.Error())
		return NewMySQLError(ER_UNKNOWN_ERROR, err.Error())
	}

	var conn *ProxyConn
	var ok bool
	for node, query := range r {
		if conn, ok = c.nodeConns[node]; !ok {
			if conn, err = node.PopConn(); err != nil {
				log.Error("node %s pop conn error %s", node.name, err.Error())
				return err
			}

			if c.isInTrans() {
				if _, err = conn.Begin(); err != nil {
					log.Error("node %s write begin error %s", node.name, err.Error())
					return err
				}
			}

			c.nodeConns[node] = conn
		}

		if err = conn.WriteCommandBuf(COM_QUERY, query); err != nil {
			log.Error("node %s write command error %s", node.name, err.Error())
			return err
		}
	}

	return nil
}

func (c *ClientConn) handleSelect(data []byte) (err error) {
	if err = c.routeQuery(data); err != nil {
		return
	}

	var result *TextResultPacket = nil

	for node, conn := range c.nodeConns {
		if r, err1 := conn.ReadTextResult(); err1 != nil {
			err = err1
			log.Error("node %s read text result error %s", node.name, err.Error())
		} else {
			if result == nil {
				result = r
			} else {
				//todo check columns defs same

				result.Rows = append(result.Rows, r.Rows...)
			}
		}
	}

	if !c.isInTrans() {
		c.clearNodeConns()
	}

	if err != nil {
		return
	} else {
		c.writeTextResult(result)
	}

	return
}

func (c *ClientConn) writeTextResult(result *TextResultPacket) error {
	count := PutLengthEncodedInt(uint64(len(result.ColumnDefs)))

	data := make([]byte, 4, 1024)
	data = append(data, count...)
	if err := c.WritePacket(data); err != nil {
		return err
	}

	for _, column := range result.ColumnDefs {
		data = data[0:4]
		data = append(data, column...)
		if err := c.WritePacket(data); err != nil {
			return err
		}
	}

	if err := c.WriteEOF(&EOFPacket{Status: c.status}); err != nil {
		return err
	}

	for _, row := range result.Rows {
		data = data[0:4]
		data = append(data, row...)
		if err := c.WritePacket(data); err != nil {
			return err
		}
	}

	if err := c.WriteEOF(&EOFPacket{Status: c.status}); err != nil {
		return err
	}

	return nil
}

func (c *ClientConn) handleExec(data []byte) (err error) {
	if err = c.routeQuery(data); err != nil {
		return
	}

	pkg := &OKPacket{Status: c.status}

	for node, conn := range c.nodeConns {
		if p, err1 := conn.ReadOK(); err1 != nil {
			err = err1
			log.Error("node %s read ok error %s", node.name, err.Error())
		} else {
			pkg.AffectedRows += p.AffectedRows
			if pkg.LastInsertId < p.LastInsertId {
				pkg.LastInsertId = p.LastInsertId
			}

			pkg.Status |= p.Status
			//now we skip warning and info
			//pkg.Warnings += p.Warnings
			//pkg.Info = p.Info
		}
	}

	if !c.isInTrans() {
		c.clearNodeConns()
	}

	if err != nil {
		return
	} else {
		c.WriteOK(pkg)
	}

	return
}

func (c *ClientConn) handleBegin() error {
	c.status |= SERVER_STATUS_IN_TRANS

	c.WriteOK(&OKPacket{Status: c.status})

	return nil
}

func (c *ClientConn) handleCommit() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for n, v := range c.nodeConns {
		if _, err1 := v.Commit(); err1 != nil {
			err = err1
			log.Error("%s commit error %s", n.name, err.Error())
		}
	}

	c.clearNodeConns()

	if err != nil {
		return
	} else {
		c.WriteOK(&OKPacket{Status: c.status})
	}

	return
}

func (c *ClientConn) handleRollback() (err error) {
	c.status &= ^SERVER_STATUS_IN_TRANS

	for n, v := range c.nodeConns {
		if _, err1 := v.Rollback(); err1 != nil {
			err = err1
			log.Error("%s rollback error %s", n.name, err.Error())
		}
	}

	c.clearNodeConns()

	if err != nil {
		return
	} else {
		c.WriteOK(&OKPacket{Status: c.status})
	}

	return
}
