package sqlparser

import (
	"github.com/siddontang/mixer/router"
	"testing"
)

/*
   range:
   node1: (-inf, 10000)
   node2: [10000, 20000)
   node3: [20000, +inf]
*/

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
    key: id
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

func checkSharding(t *testing.T, sql string, checkNodes []string) {
	r := newTestDBRule()

	ns, err := GetShardList(sql, nil, r)
	if err != nil {
		t.Fatal(err)
	} else if len(ns) != len(checkNodes) {
		t.Fatal(len(ns), len(checkNodes))
	} else {
		for i := range ns {
			if ns[i] != checkNodes[i] {
				t.Fatal(ns[i], checkNodes[i])
			}
		}
	}
}

func TestConditionSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = 5"
	checkSharding(t, sql, []string{"node6"})

	sql = "select * from test1 where id in (5, 6)"
	checkSharding(t, sql, []string{"node6", "node7"})

	sql = "select * from test1 where id > 5"
	checkSharding(t, sql, []string{"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8", "node9", "node10"})

	sql = "select * from test2 where id = 10000"
	checkSharding(t, sql, []string{"node2"})

	sql = "select * from test2 where id between 10000 and 100000"
	checkSharding(t, sql, []string{"node2", "node3"})

	sql = "select * from test2 where id > 10000"
	checkSharding(t, sql, []string{"node2", "node3"})

	sql = "select * from test2 where id >= 10000"
	checkSharding(t, sql, []string{"node2", "node3"})

	sql = "select * from test2 where id <= 10000"
	checkSharding(t, sql, []string{"node1", "node2"})
}

func TestValueSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (5)"
	checkSharding(t, sql, []string{"node6"})

	sql = "insert into test2 (id) values (10000)"
	checkSharding(t, sql, []string{"node2"})

	sql = "insert into test2 (id) values (20000)"
	checkSharding(t, sql, []string{"node3"})

	sql = "insert into test2 (id) values (200000)"
	checkSharding(t, sql, []string{"node3"})
}
