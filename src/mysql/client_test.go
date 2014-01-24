package mysql

import (
	"github.com/siddontang/golib/log"
	"sync"
	"testing"
)

var testClient *Client
var testClientOnce sync.Once

func newTestClient() *Client {
	f := func() {
		c := NewClient()

		if err := c.Connect("10.20.135.213:3306", "qing", "admin", "mixer"); err != nil {
			log.Error("%s", err.Error())
		}

		if _, err := c.Exec("set autocommit = 1"); err != nil {
			log.Error("set autocommit error %s", err.Error())
			c.Close()
		}

		testClient = c
	}

	testClientOnce.Do(f)

	return testClient
}

func TestClient_Connect(t *testing.T) {
	newTestClient()
}

func TestClient_Ping(t *testing.T) {
	c := newTestClient()

	if err := c.Ping(); err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateTable(t *testing.T) {
	s := `CREATE TABLE IF NOT EXISTS mixer_test (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	c := newTestClient()

	if _, err := c.Exec(s); err != nil {
		t.Fatal(err)
	}
}

func TestClient_Insert(t *testing.T) {
	s := `insert into mixer_test (id, str, f, e) values(1, "a", 3.14, "test1")`

	c := newTestClient()

	if pkg, err := c.Exec(s); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}
}

func TestClient_Select(t *testing.T) {
	s := `select str, f, e from mixer_test where id = 1`

	c := newTestClient()

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

func TestClient_FieldList(t *testing.T) {
	c := newTestClient()

	if result, err := c.FieldList("mixer_test", "st%"); err != nil {
		t.Fatal(err)
	} else {
		if len(result) != 1 {
			t.Fatal(len(result))
		}
	}
}

func TestClient_DeleteTable(t *testing.T) {
	c := newTestClient()

	if _, err := c.Exec("drop table mixer_test"); err != nil {
		t.Fatal(err)
	}
}
