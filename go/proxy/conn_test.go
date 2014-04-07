package proxy

import (
	. "github.com/siddontang/mixer/go/mysql"
	"testing"
)

func TestConn_Handshake(t *testing.T) {
	db := newTestDB()

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}

func TestConn_DeleteTable(t *testing.T) {
	server := newTestServer()
	nodes := server.nodes
	for _, n := range nodes {
		if _, err := n.db.Exec(`drop table if exists mixer_test_proxy_conn`); err != nil {
			t.Fatal(err)
		}
	}
}

func TestConn_CreateTable(t *testing.T) {
	s := `CREATE TABLE IF NOT EXISTS mixer_test_proxy_conn (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          u tinyint unsigned,
          i tinyint,
          ni tinyint,
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	server := newTestServer()
	nodes := server.nodes
	for _, n := range nodes {
		if _, err := n.db.Exec(s); err != nil {
			t.Fatal(err)
		}
	}
}

func TestConn_Insert(t *testing.T) {
	s := `insert into mixer_test_proxy_conn (id, str, f, e, u, i) values(1, "abc", 3.14, "test1", 255, -127)`

	db := newTestDB()
	if r, err := db.Exec(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 1 {
			t.Fatal(r.AffectedRows)
		}
	}
}

func TestConn_Select(t *testing.T) {
	s := `select str, f, e, u, i, ni from mixer_test_proxy_conn where id = 1`

	db := newTestDB()
	if r, err := db.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if r.RowNumber() != 1 {
			t.Fatal(r.RowNumber())
		}

		if r.ColumnNumber() != 6 {
			t.Fatal(r.ColumnNumber())
		}

		if v, _ := r.GetString(0, 0); v != `abc` {
			t.Fatal(v)
		}

		if v, _ := r.GetFloat(0, 1); v != 3.14 {
			t.Fatal(v)
		}

		if v, _ := r.GetString(0, 2); v != `test1` {
			t.Fatal(v)
		}

		if v, _ := r.GetUint(0, 3); v != 255 {
			t.Fatal(v)
		}

		if v, _ := r.GetInt(0, 4); v != -127 {
			t.Fatal(v)
		}

		if v, _ := r.IsNull(0, 5); !v {
			t.Fatal("ni not null")
		}
	}
}

func TestConn_Update(t *testing.T) {
	s := `update mixer_test_proxy_conn set str = "123" where id = 1`

	db := newTestDB()

	if _, err := db.Exec(s); err != nil {
		t.Fatal(err)
	}

	if r, err := db.Query(`select str from mixer_test_proxy_conn where id = 1`); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetString(0, 0); v != `123` {
			t.Fatal(v)
		}
	}
}

func TestConn_Replace(t *testing.T) {
	s := `replace into mixer_test_proxy_conn (id, str, f) values(1, "abc", 3.14159)`

	db := newTestDB()

	if r, err := db.Exec(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 2 {
			t.Fatal(r.AffectedRows)
		}
	}

	s = `replace into mixer_test_proxy_conn (id, str) values(2, "abcb")`

	if r, err := db.Exec(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 1 {
			t.Fatal(r.AffectedRows)
		}
	}

	s = `select str, f from mixer_test_proxy_conn`

	if r, err := db.Query(s); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetString(0, 0); v != `abc` {
			t.Fatal(v)
		}

		if v, _ := r.GetString(1, 0); v != `abcb` {
			t.Fatal(v)
		}

		if v, _ := r.GetFloat(0, 1); v != 3.14159 {
			t.Fatal(v)
		}

		if v, _ := r.IsNull(1, 1); !v {
			t.Fatal(v)
		}
	}
}

func TestConn_Delete(t *testing.T) {
	s := `delete from mixer_test_proxy_conn where id = 2`

	db := newTestDB()

	if r, err := db.Exec(s); err != nil {
		t.Fatal(err)
	} else {
		if r.AffectedRows != 1 {
			t.Fatal(r.AffectedRows)
		}
	}
}

func TestConn_SetAutoCommit(t *testing.T) {
	db := newTestDB()

	if r, err := db.Exec("set autocommit = 1"); err != nil {
		t.Fatal(err)
	} else {
		if !(r.Status&SERVER_STATUS_AUTOCOMMIT > 0) {
			t.Fatal(r.Status)
		}
	}

	if r, err := db.Exec("set autocommit = 0"); err != nil {
		t.Fatal(err)
	} else {
		if !(r.Status&SERVER_STATUS_AUTOCOMMIT == 0) {
			t.Fatal(r.Status)
		}
	}

	if r, err := db.Query("select 1"); err != nil {
		t.Fatal(err)
	} else {
		if !(r.Status&SERVER_STATUS_AUTOCOMMIT > 0) {
			t.Fatal(r.Status)
		}
	}
}

func TestConn_Trans(t *testing.T) {
	db := newTestDB()

	var tx1 *Tx
	var tx2 *Tx
	var err error

	if tx1, err = db.Begin(); err != nil {
		t.Fatal(err)
	}

	if tx2, err = db.Begin(); err != nil {
		t.Fatal(err)
	}

	if _, err := tx1.Exec(`insert into mixer_test_proxy_conn (id, str) values (111, "abc")`); err != nil {
		t.Fatal(err)
	}

	if r, err := tx2.Query(`select str from mixer_test_proxy_conn where id = 111`); err != nil {
		t.Fatal(err)
	} else {
		if r.RowNumber() != 0 {
			t.Fatal(r.RowNumber())
		}
	}

	if err := tx1.Commit(); err != nil {
		t.Fatal(err)
	}

	if err := tx2.Commit(); err != nil {
		t.Fatal(err)
	}

	if r, err := db.Query(`select str from mixer_test_proxy_conn where id = 111`); err != nil {
		t.Fatal(err)
	} else {
		if r.RowNumber() != 1 {
			t.Fatal(r.RowNumber())
		}

		if v, _ := r.GetString(0, 0); v != `abc` {
			t.Fatal(v)
		}
	}
}

func TestConn_SetNames(t *testing.T) {
	db := newTestDB()

	c, err := db.GetConn()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if err := c.SetCharset("gb2312"); err != nil {
		t.Fatal(err)
	}

	if r, err := c.Query("select 1 + 1"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetInt(0, 0); v != 2 {
			t.Fatal(v)
		}
	}
}

func TestConn_LastInsertId(t *testing.T) {
	db := newTestDB()

	c, err := db.GetConn()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	s := `CREATE TABLE IF NOT EXISTS mixer_test_conn_id (
          id BIGINT(64) UNSIGNED AUTO_INCREMENT NOT NULL,
          str VARCHAR(256),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	server := newTestServer()
	nodes := server.nodes
	for _, n := range nodes {
		if _, err := n.db.Exec(s); err != nil {
			t.Fatal(err)
		}
	}

	r, err := c.Exec(`insert into mixer_test_conn_id (str) values ("abc")`)
	if err != nil {
		t.Fatal(err)
	}

	lastId := r.InsertId

	if r, err := c.Query(`select last_insert_id();`); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetUint(0, 0); v != lastId {
			t.Fatal(v)
		}
	}

	if r, err := c.Query(`select last_insert_id() as a;`); err != nil {
		t.Fatal(err)
	} else {
		if string(r.Fields[0].Name) != "a" {
			t.Fatal(string(r.Fields[0].Name))
		}

		if v, _ := r.GetUint(0, 0); v != lastId {
			t.Fatal(v)
		}
	}

	for _, n := range nodes {
		if _, err := n.db.Exec(`drop table if exists mixer_test_conn_id`); err != nil {
			t.Fatal(err)
		}
	}
}

func TestConn_RowCount(t *testing.T) {
	db := newTestDB()

	c, err := db.GetConn()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	r, err := c.Exec(`insert into mixer_test_proxy_conn (id, str) values (1002, "abc")`)
	if err != nil {
		t.Fatal(err)
	}

	row := r.AffectedRows

	if r, err := c.Query("select row_count();"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetUint(0, 0); v != row {
			t.Fatal(v)
		}
	}

	if r, err := c.Query("select row_count() as b;"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetInt(0, 0); v != -1 {
			t.Fatal(v)
		}
	}
}

func TestConn_SelectVersion(t *testing.T) {
	db := newTestDB()

	c, err := db.GetConn()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if r, err := c.Query("select version();"); err != nil {
		t.Fatal(err)
	} else {
		if v, _ := r.GetString(0, 0); v != ServerVersion {
			t.Fatal(v)
		}
	}
}
