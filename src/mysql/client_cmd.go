package mysql

import (
	"github.com/siddontang/golib/log"
)

func (c *Client) Quit() {
	err := c.writeCommandPacket(COM_QUIT)
	if err != nil {
		log.Error("quit error %s", err.Error())
		return
	}

	c.Close()

	return
}

func (c *Client) UseDB(db string) error {
	err := c.writeCommandPacketStr(COM_INIT_DB, db)
	if err != nil {
		log.Error("use db error %s", err.Error())
		return err
	}

	return c.readResultOKPacket()
}

func (c *Client) Ping() error {
	err := c.writeCommandPacket(COM_PING)

	if err != nil {
		log.Error("ping error %s", err.Error())
		return err
	}

	return c.readResultOKPacket()
}

func (c *Client) Query(query string) (*ResultSet, error) {
	err := c.writeCommandPacketStr(COM_QUERY, query)

	if err != nil {
		log.Error("query error %s", err.Error())
		return nil, err
	}

	var columns [][]byte
	var rows [][]byte

	columns, rows, err = c.readTextResultSetPacket()
	if err != nil {
		log.Error("read text result set error %s", err.Error())
		return nil, err
	}

	return NewTextResultSet(columns, rows)
}

func (c *Client) Exec(query string) error {
	err := c.writeCommandPacketStr(COM_QUERY, query)

	if err != nil {
		log.Error("exec error %s", err.Error())
		return err
	}

	return c.readResultOKPacket()
}
