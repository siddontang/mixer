package client

import (
	"container/list"
	"fmt"
	. "github.com/siddontang/mixer/mysql"
	"sync"
)

type DB struct {
	sync.Mutex

	addr      string
	user      string
	password  string
	db        string
	idleConns int

	conns *list.List
}

func Open(addr string, user string, password string, dbName string) (*DB, error) {
	db := new(DB)

	db.addr = addr
	db.user = user
	db.password = password
	db.db = dbName

	db.conns = list.New()

	return db, nil
}

func (db *DB) Addr() string {
	return db.addr
}

func (db *DB) ConfigString() string {
	return fmt.Sprintf("%s:%s@%s/%s?idleConns=%v&conns=%v",
		db.user, db.password, db.addr, db.db, db.idleConns, db.conns.Len())
}

func (db *DB) Close() error {
	db.Lock()

	for {
		if db.conns.Len() > 0 {
			v := db.conns.Back()
			co := v.Value.(*Conn)
			db.conns.Remove(v)

			co.Close()

		} else {
			break
		}
	}

	db.Unlock()

	return nil
}

func (db *DB) Ping() error {
	c, err := db.PopConn()
	if err != nil {
		return err
	}

	err = c.Ping()
	db.PushConn(c, err)
	return err
}

func (db *DB) SetIdleConns(num int) {
	db.idleConns = num
}

func (db *DB) newConn() (*Conn, error) {
	co := new(Conn)

	if err := co.Connect(db.addr, db.user, db.password, db.db); err != nil {
		return nil, err
	}

	return co, nil
}

func (db *DB) tryReuse(co *Conn) error {
	if co.IsInTransaction() {
		//we can not reuse a connection in transaction status
		if err := co.Rollback(); err != nil {
			return err
		}
	}

	if !co.IsAutoCommit() {
		//we can not  reuse a connection not in autocomit
		if _, err := co.exec("set autocommit = 1"); err != nil {
			return err
		}
	}

	//connection may be set names early
	//we must use default utf8
	if co.GetCharset() != DEFAULT_CHARSET {
		if err := co.SetCharset(DEFAULT_CHARSET); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) PopConn() (co *Conn, err error) {
	db.Lock()
	if db.conns.Len() > 0 {
		v := db.conns.Front()
		co = v.Value.(*Conn)
		db.conns.Remove(v)
	}
	db.Unlock()

	if co != nil {
		if err := co.Ping(); err == nil {
			if err := db.tryReuse(co); err == nil {
				//connection may alive
				return co, nil
			}
		}
		co.Close()
	}

	return db.newConn()
}

func (db *DB) PushConn(co *Conn, err error) {
	var closeConn *Conn = nil

	if err != nil {
		closeConn = co
	} else {
		if db.idleConns > 0 {
			db.Lock()

			if db.conns.Len() >= db.idleConns {
				v := db.conns.Front()
				closeConn = v.Value.(*Conn)
				db.conns.Remove(v)
			}

			db.conns.PushBack(co)

			db.Unlock()

		} else {
			closeConn = co
		}

	}

	if closeConn != nil {
		closeConn.Close()
	}
}

type SqlConn struct {
	*Conn

	db *DB
}

func (p *SqlConn) Close() {
	if p.Conn != nil {
		p.db.PushConn(p.Conn, p.Conn.pkgErr)
		p.Conn = nil
	}
}

func (db *DB) GetConn() (*SqlConn, error) {
	c, err := db.PopConn()
	return &SqlConn{c, db}, err
}
