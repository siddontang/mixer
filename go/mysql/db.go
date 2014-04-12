package mysql

import (
	"container/list"
	"encoding/json"
	"sync"
)

type Config struct {
	Addr      string `json:"addr"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DB        string `json:"db"`
	IdleConns int    `json:"idle_conns"`
}

type DB struct {
	sync.Mutex

	cfg   *Config
	conns *list.List
}

func NewDB(jsonConfig json.RawMessage) (*DB, error) {
	cfg := new(Config)

	err := json.Unmarshal(jsonConfig, cfg)
	if err != nil {
		return nil, err
	}

	db := new(DB)
	db.cfg = cfg
	db.conns = list.New()

	return db, nil
}

func (db *DB) Addr() string {
	return db.cfg.Addr
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

func (db *DB) newConn() (*Conn, error) {
	co := new(Conn)

	if err := co.Connect(db.cfg.Addr, db.cfg.User, db.cfg.Password, db.cfg.DB); err != nil {
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

	if err == ErrBadConn {
		closeConn = co
	} else {
		if db.cfg.IdleConns > 0 {
			db.Lock()

			if db.conns.Len() >= db.cfg.IdleConns {
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
	p.db.PushConn(p.Conn, nil)
}

func NewSqlConn(db *DB) (*SqlConn, error) {
	c, err := db.PopConn()
	return &SqlConn{c, db}, err
}
