package sqlparser

import (
	"testing"
)

func TestSet(t *testing.T) {
	sql := "set names gbk"

	_, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSimpleSelect(t *testing.T) {
	sql := "select last_insert_id() as a"

	_, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdmin(t *testing.T) {
	sql := `admin upnode("node1", "master", "127.0.0.1")`

	_, err := Parse(sql)

	if err != nil {
		t.Fatal(err)
	}
}
