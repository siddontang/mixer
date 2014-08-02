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
