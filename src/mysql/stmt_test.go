package mysql

import (
	"testing"
)

func TestStmt_CreateTable(t *testing.T) {
	str := `CREATE TABLE IF NOT EXISTS mixer_test_stmt (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err = s.Exec(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_Delete(t *testing.T) {
	str := `delete from mixer_test_stmt`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Exec(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_Insert(t *testing.T) {
	str := `insert into mixer_test_stmt (id, str, f, e) values (?, ?, ?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Exec(1, "a", 3.14, "test1"); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()
}

func TestStmt_Select(t *testing.T) {
	str := `select str, f, e from mixer_test_stmt where id = ?`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if result, err := s.Query(1); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Data) != 1 {
			t.Fatal(len(result.Data))
		}

		if len(result.Fields) != 3 {
			t.Fatal(len(result.Fields))
		}

		if str, _ := result.GetString(0, 0); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloat(0, 1); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetString(0, 2); e != "test1" {
			t.Fatal("invalid e", e)
		}

		if str, _ := result.GetStringByName(0, "str"); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloatByName(0, "f"); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetStringByName(0, "e"); e != "test1" {
			t.Fatal("invalid e", e)
		}

	}

	s.Close()
}

func TestStmt_DropTable(t *testing.T) {
	str := `drop table mixer_test_stmt`

	c := newTestConn()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Exec(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}
