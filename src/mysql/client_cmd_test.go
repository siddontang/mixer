package mysql

import (
	"testing"
)

func TestClient_InitDB(t *testing.T) {
	client := newTestClient()

	err := client.UseDB("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Ping(t *testing.T) {
	client := newTestClient()

	err := client.Ping()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Select(t *testing.T) {
	client := newTestClient()

	r, err := client.Query("select 1 + 1 as a")
	if err != nil {
		t.Fatal(err)
	}

	var d int64
	d, err = r.GetInt(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	if d != 2 {
		t.Fatal("invalid data: ", d)
	}
}

func TestClient_Create(t *testing.T) {
	client := newTestClient()

	s := `CREATE TABLE IF NOT EXISTS mixer_test (
          id BIGINT(64) UNSIGNED  NOT NULL,
          ts TIMESTAMP, 
          y YEAR, 
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	err := client.Exec(s)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Insert(t *testing.T) {
	client := newTestClient()

	s := `insert into mixer_test (id, ts, y, str, f, e) values (0, "2014-01-07 10:00:00", 2014, "hello world", "3.14", "test1")`

	err := client.Exec(s)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(client.affectedRows, client.lastInsertId)

	if client.affectedRows != 1 {
		t.Fatal("affectedRows: ", client.affectedRows)
	}

	if client.lastInsertId != 0 {
		t.Fatal("last insert id: ", client.lastInsertId)
	}
}

func TestClient_Query(t *testing.T) {
	client := newTestClient()

	s := `select ts, y, str, f, e from mixer_test where id = 0`

	r, err := client.Query(s)

	if err != nil {
		t.Fatal(err)
	}

	var ts string
	ts, err = r.GetString(0, 0)

	if err != nil {
		t.Fatal(err)
	}

	if ts != "2014-01-07 10:00:00" {
		t.Fatal(ts)
	}

	var y int64

	y, err = r.GetInt(0, 1)

	if err != nil {
		t.Fatal(err)
	}

	if y != 2014 {
		t.Fatal(y)
	}

	var str string

	str, err = r.GetStringByName(0, "str")

	if err != nil {
		t.Fatal(err)
	}

	if str != "hello world" {
		t.Fatal(str)
	}

	var e string

	e, err = r.GetStringByName(0, "e")

	if err != nil {
		t.Fatal(err)
	}

	if e != "test1" {
		t.Fatal(e)
	}

	var f float64

	f, err = r.GetFloatByName(0, "f")

	if err != nil {
		t.Fatal(err)
	}

	if f != 3.14 {
		t.Fatal(f)
	}
}

func TestClient_Drop(t *testing.T) {
	client := newTestClient()

	s := `drop table mixer_test`

	err := client.Exec(s)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Quit(t *testing.T) {
	client := newTestClient()

	client.Quit()

	err := client.ReConnect()
	if err != nil {
		t.Fatal(err)
	}
}
