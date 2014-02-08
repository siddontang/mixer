package mysql

import (
	"container/list"
	"github.com/siddontang/golib/log"
	"sync"
)

type Client struct {
	addr     string
	user     string
	password string
	db       string

	maxIdleConns int

	lock  sync.Mutex
	conns *list.List
}

func NewClient(addr string, user string, password string, db string, maxIdleConns int) *Client {
	c := new(Client)

	c.addr = addr
	c.user = user
	c.password = password
	c.db = db
	c.maxIdleConns = maxIdleConns

	c.conns = list.New()

	return c
}

func (c *Client) popConn() (*conn, error) {
	var co *conn

	c.lock.Lock()
	if v := c.conns.Back(); v != nil {
		c.conns.Remove(v)
		co = v.Value.(*conn)
	}
	c.lock.Unlock()

	if co != nil {
		if err := co.Ping(); err == nil {
			//connection has alive
			return co, nil
		}
	}

	co = new(conn)

	if err := co.Connect(c.addr, c.user, c.password, c.db); err != nil {
		log.Error("connect %s error %s", c.addr, err.Error())
		return nil, err
	}

	//we must always use autocommit
	if _, err := co.Exec("set autocommit = 1"); err != nil {
		log.Error("set autocommit error %s", err.Error())
		co.Close()

		return nil, err
	}

	return co, nil
}

func (c *Client) pushConn(co *conn) {
	var closeConn *conn
	c.lock.Lock()

	if c.conns.Len() > c.maxIdleConns {
		oldConn := c.conns.Front()
		c.conns.Remove(oldConn)

		closeConn = oldConn.Value.(*conn)
	}

	c.conns.PushBack(co)

	c.lock.Unlock()

	if closeConn != nil {
		closeConn.Close()
	}
}

func (c *Client) Get() (Conn, error) {
	conn, err := c.popConn()
	if err != nil {
		return nil, err
	}

	return &poolConn{c, conn}, nil
}

func (c *Client) Exec(query string) (*OKPacket, error) {
	conn, err := c.Get()
	if err != nil {
		return nil, err
	}

	var r *OKPacket
	r, err = conn.Exec(query)
	conn.Close()
	return r, err
}

func (c *Client) Query(query string) (*Resultset, error) {
	conn, err := c.Get()
	if err != nil {
		return nil, err
	}

	var r *Resultset
	r, err = conn.Query(query)
	conn.Close()
	return r, err
}

type poolConn struct {
	client *Client
	*conn
}

func (c *poolConn) Close() error {
	c.client.pushConn(c.conn)
	return nil
}
