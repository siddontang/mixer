package proxy

import (
	"github.com/siddontang/golib/log"
	"mysql"
	"sync"
	"testing"
)

func newTestClient() *mysql.Conn {
	c := mysql.NewConn()

	if err := c.Connect("127.0.0.1:3306", "qing", "admin", "mixer"); err != nil {
		return nil
	}

	return c
}

var testMySQLConn *mysql.MySQLConn
var testMySQLConnOnce sync.Once

func newTestMySQLConn() *mysql.Conn {
	f := func() {
		c := new(msyql.MySQLConn)

		if err := c.Connect("10.20.135.213:3306", "qing", "admin", "mixer"); err != nil {
			log.Error("%s", err.Error())
		}

		if _, err := c.Exec("set autocommit = 1"); err != nil {
			log.Error("set autocommit error %s", err.Error())
			c.Close()
		}

		testMySQLConn = c
	}

	testMySQLConnOnce.Do(f)

	return testMySQLConn
}

func TestClientConn_Handshake(t *testing.T) {
	newTestServer()

	c := newTestClient()
	if c == nil {
		t.Fatal("connect failed")
	}

	c.Close()
}

func TestClientConn_CreateTable(t *testing.T) {
	s := `CREATE TABLE IF NOT EXISTS mixer_test_clientconn (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	c := newTestMySQLConn()

	if _, err := c.Exec(s); err != nil {
		t.Fatal(err)
	}

}

func TestClientConn_Delete(t *testing.T) {
	s := `delete from mixer_test_clientconn`

	c := newTestClient()
	defer c.Close()

	_, err := c.Exec(s)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientConn_Insert(t *testing.T) {
	s := `insert into mixer_test_clientconn (id, str, f, e) values (1, "abc", 3.14, "test1")`

	c := newTestClient()
	defer c.Close()

	pkg, err := c.Exec(s)
	if err != nil {
		t.Fatal(err)
	}

	if pkg.AffectedRows != 1 {
		t.Fatal(pkg.AffectedRows)
	}
}

func TestClientConn_Select(t *testing.T) {
	s := `select str, f, e from mixer_test_clientconn where id = 1`

	c := newTestClient()
	defer c.Close()

	result, err := c.Query(s)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.ColumnDefs) != 3 {
		t.Fatal(len(result.ColumnDefs))
	}

	if len(result.Rows) != 1 {
		t.Fatal(len(result.Rows))
	}
}

func TestClientConn_Rollback(t *testing.T) {
	c := newTestClient()
	defer c.Close()

	if _, err := c.Begin(); err != nil {
		t.Fatal(err)
	}

	s := `insert into mixer_test_clientconn (id, str, f, e) values (2, "abc", 3.14, "test1")`
	if _, err := c.Exec(s); err != nil {
		t.Fatal(err)
	}

	c1 := newTestClient()
	defer c1.Close()

	s = `select id from mixer_test_clientconn`

	if result, err := c1.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Rows) != 1 {
			t.Fatal(len(result.Rows))
		}
	}

	if _, err := c.Rollback(); err != nil {
		t.Fatal(err)
	}

	if result, err := c1.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Rows) != 1 {
			t.Fatal(len(result.Rows))
		}
	}
}

func TestClientConn_Commit(t *testing.T) {
	c := newTestClient()
	defer c.Close()

	if _, err := c.Begin(); err != nil {
		t.Fatal(err)
	}

	s := `insert into mixer_test_clientconn (id, str, f, e) values (2, "abc", 3.14, "test1")`
	if _, err := c.Exec(s); err != nil {
		t.Fatal(err)
	}

	c1 := newTestClient()
	defer c1.Close()

	s = `select id from mixer_test_clientconn`

	if result, err := c1.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Rows) != 1 {
			t.Fatal(len(result.Rows))
		}
	}

	if _, err := c.Commit(); err != nil {
		t.Fatal(err)
	}

	if result, err := c1.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Rows) != 2 {
			t.Fatal(len(result.Rows))
		}
	}
}

func TestClientConn_DeleteTable(t *testing.T) {
	c := newTestMySQLConn()

	if _, err := c.Exec("drop table mixer_test_clientconn"); err != nil {
		t.Fatal(err)
	}
}
