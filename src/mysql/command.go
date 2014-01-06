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
