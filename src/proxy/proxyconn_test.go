package proxy

import (
	"github.com/siddontang/golib/log"
	"sync"
	"testing"
)

var testProxyConn *ProxyConn
var testProxyConnOnce sync.Once

func newTestProxyConn() *ProxyConn {
	f := func() {
		c := NewProxyConn()

		if err := c.Connect("10.20.135.213:3306", "qing", "admin", "mixer"); err != nil {
			log.Error("%s", err.Error())
		}

		if _, err := c.Exec("set autocommit = 1"); err != nil {
			log.Error("set autocommit error %s", err.Error())
			c.Close()
		}

		testProxyConn = c
	}

	testProxyConnOnce.Do(f)

	return testProxyConn
}

func TestProxyConn_Connect(t *testing.T) {
	newTestProxyConn()
}

func TestProxyConn_Ping(t *testing.T) {
	c := newTestProxyConn()

	if err := c.Ping(); err != nil {
		t.Fatal(err)
	}
}

func TestProxyConn_CreateTable(t *testing.T) {
	s := `CREATE TABLE IF NOT EXISTS mixer_test_proxyconn (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	c := newTestProxyConn()

	if _, err := c.Exec(s); err != nil {
		t.Fatal(err)
	}
}

func TestProxyConn_Insert(t *testing.T) {
	s := `insert into mixer_test_proxyconn (id, str, f, e) values(1, "a", 3.14, "test1")`

	c := newTestProxyConn()

	if pkg, err := c.Exec(s); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}
}

func TestProxyConn_Select(t *testing.T) {
	s := `select str, f, e from mixer_test_proxyconn where id = 1`

	c := newTestProxyConn()

	if result, err := c.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if len(result.ColumnDefs) != 3 {
			t.Fatal(len(result.ColumnDefs))
		}

		if len(result.Rows) != 1 {
			t.Fatal(len(result.Rows))
		}
	}
}

func TestProxyConn_FieldList(t *testing.T) {
	c := newTestProxyConn()

	if result, err := c.FieldList("mixer_test_proxyconn", "st%"); err != nil {
		t.Fatal(err)
	} else {
		if len(result) != 1 {
			t.Fatal(len(result))
		}
	}
}

func TestProxyConn_DeleteTable(t *testing.T) {
	c := newTestProxyConn()

	if _, err := c.Exec("drop table mixer_test_proxyconn"); err != nil {
		t.Fatal(err)
	}
}
