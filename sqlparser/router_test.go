package sqlparser

import (
	"github.com/siddontang/mixer/router"
	"testing"
)

func newTestDBRule() *router.DBRules {
	var s = `
rules:
-
    db: mixer
    table: test1
    key: id
    type: hash
    nodes: node(1-10)

-
    db: mixer
    table: test2 
    key: uid
    nodes: node1,node2,node3    
    type: range
    # range is -inf-10000 10000-20000 20000-+inf 
    range: -0000000000002710-0000000000004e20-

-   db: mixer
    table: 
    key:
    nodes: node1
    type: default
`

	r, err := router.NewRouterConfigData([]byte(s))
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return r.GetDBRules("mixer")
}

func TestSelectSharding(t *testing.T) {
	r := newTestDBRule()

	var sql string

	sql = "select * from test1 where id = 5"

	ns, err := GetShardList(sql, nil, r)
	if err != nil {
		t.Fatal(err)
	} else if len(ns) != 1 {
		t.Fatal(len(ns))
	} else if ns[0] != "node6" {
		t.Fatal(ns[0])
	}
}
