package sqlparser

import (
	"github.com/siddontang/mixer/config"
	"github.com/siddontang/mixer/router"

	"fmt"
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

	cfg, err := config.ParseConfigData([]byte(s))
	if err != nil {
		println(err.Error())
		panic(err)
	}

	var r *router.Router

	r, err = router.NewRouter(cfg)
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return r.GetDBRules("mixer")
}

func checkSharding(t *testing.T, sql string, checkNodeIndex ...int) {
	r := newTestDBRule()

	ns, err := GetShardListIndex(sql, r)
	if err != nil {
		t.Fatal(sql, err)
	} else if len(ns) != len(checkNodeIndex) {
		s := fmt.Sprintf("%v %v", ns, checkNodeIndex)
		t.Fatal(sql, s)
	} else {

		for i := range ns {
			if ns[i] != checkNodeIndex[i] {
				s := fmt.Sprintf("%v %v", ns, checkNodeIndex)
				t.Fatal(sql, s, i)
			}
		}
	}
}

func TestConditionSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = 5"
	checkSharding(t, sql, 5)

	sql = "select * from test1 where id in (5, 6)"
	checkSharding(t, sql, 5, 6)

	sql = "select * from test1 where id > 5"
	checkSharding(t, sql, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) and id in (5, 6, 7)"
	checkSharding(t, sql, 5, 6)

	sql = "select * from test1 where id in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (5, 6)"
	checkSharding(t, sql, 0, 1, 2, 3, 4, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) or (id in (5, 6, 7,8) and id in (1,5,7))"
	checkSharding(t, sql, 5, 6, 7)

	sql = "select * from test2 where id = 10000"
	checkSharding(t, sql, 1)

	sql = "select * from test2 where id between 10000 and 100000"
	checkSharding(t, sql, 1, 2)

	sql = "select * from test2 where id not between 1000 and 100000"
	checkSharding(t, sql, 0, 2)

	sql = "select * from test2 where id not between 10000 and 100000"
	checkSharding(t, sql, 0, 2)

	sql = "select * from test2 where id > 10000"
	checkSharding(t, sql, 1, 2)

	sql = "select * from test2 where id >= 10000"
	checkSharding(t, sql, 1, 2)

	sql = "select * from test2 where id <= 10000"
	checkSharding(t, sql, 0, 1)

	sql = "select * from test2 where id < 10000"
	checkSharding(t, sql, 0)

	sql = "select * from test2 where  10000 < id"
	checkSharding(t, sql, 1, 2)

	sql = "select * from test2 where  10000 <= id"
	checkSharding(t, sql, 1, 2)

	sql = "select * from test2 where  10000 > id"
	checkSharding(t, sql, 0)

	sql = "select * from test2 where  10000 >= id"
	checkSharding(t, sql, 0, 1)

	sql = "select * from test2 where id >= 10000 and id <= 100000"
	checkSharding(t, sql, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or id < 100"
	checkSharding(t, sql, 0, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or (id < 100 and name > 100000)"
	checkSharding(t, sql, 0, 1, 2)

	sql = "select * from test2 where id in (1, 10000)"
	checkSharding(t, sql, 0, 1)

	sql = "select * from test2 where id not in (1, 10000)"
	checkSharding(t, sql, 0, 1, 2)

	sql = "select * from test2 where id in (1000, 10000)"
	checkSharding(t, sql, 0, 1)

	sql = "select * from test2 where id > -1"
	checkSharding(t, sql, 0, 1, 2)
}

func TestValueSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (5)"
	checkSharding(t, sql, 5)

	sql = "insert into test2 (id) values (10000)"
	checkSharding(t, sql, 1)

	sql = "insert into test2 (id) values (20000)"
	checkSharding(t, sql, 2)

	sql = "insert into test2 (id) values (200000)"
	checkSharding(t, sql, 2)
}

func testCheckList(t *testing.T, l []int, checkList ...int) {
	if len(l) != len(checkList) {
		t.Fatal("invalid list len", len(l), len(checkList))
	}

	for i := 0; i < len(l); i++ {
		if l[i] != checkList[i] {
			t.Fatal("invalid list item", l[i], i)
		}
	}
}

func TestListSet(t *testing.T) {
	var l1 []int
	var l2 []int
	var l3 []int

	l1 = []int{1, 2, 3}
	l2 = []int{2}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 3}
	l2 = []int{2, 3}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2, 3)

	l1 = []int{1, 2, 4}
	l2 = []int{2, 3}

	l3 = interList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 4}
	l2 = []int{}

	l3 = interList(l1, l2)
	testCheckList(t, l3)

	l1 = []int{1, 2, 3}
	l2 = []int{2}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3)

	l1 = []int{1, 2, 4}
	l2 = []int{3}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{2, 3, 4}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{}

	l3 = unionList(l1, l2)
	testCheckList(t, l3, 1, 2, 3)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{2}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 3, 4)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 2, 3, 4)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{1, 3, 5}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 2, 4)

	l1 = []int{1, 2, 3}
	l2 = []int{1, 3, 5, 6}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 2)

	l1 = []int{1, 2, 3, 4}
	l2 = []int{2, 3}

	l3 = differentList(l1, l2)
	testCheckList(t, l3, 1, 4)
}
