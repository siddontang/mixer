package sqlparser

import (
	"github.com/siddontang/mixer/config"
	"github.com/siddontang/mixer/router"

	"fmt"
	"testing"
)

func newTestDBRule() *router.Router {
	var s = `
schemas :
-
  db : mixer 
  nodes: [node1,node2,node3,node4,node5,node6,node7,node8,node9,node10]
  rules:
    default: node1
    shard:
      -   
        table: test1
        key: id
        nodes: [node1,node2,node3,node4,node5,node6,node7,node8,node9,node10]
        type: hash

      -   
        table: test2
        key: id
        type: range
        nodes: [node1,node2,node3]
        range: -10000-20000-
`

	cfg, err := config.ParseConfigData([]byte(s))
	if err != nil {
		println(err.Error())
		panic(err)
	}

	var r *router.Router

	r, err = router.NewRouter(&cfg.Schemas[0])
	if err != nil {
		println(err.Error())
		panic(err)
	}

	return r
}

func checkSharding(t *testing.T, sql string, args []int, checkNodeIndex ...int) {
	r := newTestDBRule()

	bindVars := make(map[string]interface{}, len(args))
	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}
	ns, err := GetShardListIndex(sql, r, bindVars)
	if err != nil {
		t.Fatal(sql, err)
	} else if len(ns) != len(checkNodeIndex) {
		s := fmt.Sprintf("%v %v", ns, checkNodeIndex)
		t.Fatal(sql, s)
	} else {
		for i := range ns {
			if ns[i] != checkNodeIndex[i] {
				s := fmt.Sprintf("%v %v", ns, checkNodeIndex)
				panic(sql)
				t.Fatal(sql, s, i)
			}
		}
	}
}

func TestConditionSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = 5"
	checkSharding(t, sql, nil, 5)

	sql = "select * from test1 where id in (5, 6)"
	checkSharding(t, sql, nil, 5, 6)

	sql = "select * from test1 where id > 5"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) and id in (5, 6, 7)"
	checkSharding(t, sql, nil, 5, 6)

	sql = "select * from test1 where id in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, nil, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (5, 6) or id in (5, 6, 7,8)"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (5, 6)"
	checkSharding(t, sql, nil, 0, 1, 2, 3, 4, 7, 8, 9)

	sql = "select * from test1 where id in (5, 6) or (id in (5, 6, 7,8) and id in (1,5,7))"
	checkSharding(t, sql, nil, 5, 6, 7)

	sql = "select * from test2 where id = 10000"
	checkSharding(t, sql, nil, 1)

	sql = "select * from test2 where id between 10000 and 100000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where id not between 1000 and 100000"
	checkSharding(t, sql, nil, 0, 2)

	sql = "select * from test2 where id not between 10000 and 100000"
	checkSharding(t, sql, nil, 0, 2)

	sql = "select * from test2 where id > 10000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where id >= 10000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where id <= 10000"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id < 10000"
	checkSharding(t, sql, nil, 0)

	sql = "select * from test2 where  10000 < id"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where  10000 <= id"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where  10000 > id"
	checkSharding(t, sql, nil, 0)

	sql = "select * from test2 where  10000 >= id"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id >= 10000 and id <= 100000"
	checkSharding(t, sql, nil, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or id < 100"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where (id >= 10000 and id <= 100000) or (id < 100 and name > 100000)"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id in (1, 10000)"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id not in (1, 10000)"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id in (1000, 10000)"
	checkSharding(t, sql, nil, 0, 1)

	sql = "select * from test2 where id > -1"
	checkSharding(t, sql, nil, 0, 1, 2)

	sql = "select * from test2 where id > -1 and id < 11000"
	checkSharding(t, sql, nil, 0, 1)
}

func TestConditionVarArgSharding(t *testing.T) {
	var sql string

	sql = "select * from test1 where id = ?"
	checkSharding(t, sql, []int{5}, 5)

	sql = "select * from test1 where id in (?, ?)"
	checkSharding(t, sql, []int{5, 6}, 5, 6)

	sql = "select * from test1 where id > ?"
	checkSharding(t, sql, []int{5}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id in (?, ?) and id in (?, ?, ?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7}, 5, 6)

	sql = "select * from test1 where id in (?, ?) or id in (?, ?,?,?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8}, 5, 6, 7, 8)

	sql = "select * from test1 where id not in (?, ?) or id in (?, ?, ?,?)"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8}, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	sql = "select * from test1 where id not in (?, ?)"
	checkSharding(t, sql, []int{5, 6}, 0, 1, 2, 3, 4, 7, 8, 9)

	sql = "select * from test1 where id in (?, ?) or (id in (?, ?, ?,?) and id in (?,?,?))"
	checkSharding(t, sql, []int{5, 6, 5, 6, 7, 8, 1, 5, 7}, 5, 6, 7)

	sql = "select * from test2 where id = ?"
	checkSharding(t, sql, []int{10000}, 1)

	sql = "select * from test2 where id between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 1, 2)

	sql = "select * from test2 where id not between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 2)

	sql = "select * from test2 where id not between ? and ?"
	checkSharding(t, sql, []int{10000, 100000}, 0, 2)

	sql = "select * from test2 where id > ?"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where id >= ?"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where id <= ?"
	checkSharding(t, sql, []int{10000}, 0, 1)

	sql = "select * from test2 where id < ?"
	checkSharding(t, sql, []int{10000}, 0)

	sql = "select * from test2 where  ? < id"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where  ? <= id"
	checkSharding(t, sql, []int{10000}, 1, 2)

	sql = "select * from test2 where  ? > id"
	checkSharding(t, sql, []int{10000}, 0)

	sql = "select * from test2 where  ? >= id"
	checkSharding(t, sql, []int{10000}, 0, 1)

	sql = "select * from test2 where id >= ? and id <= ?"
	checkSharding(t, sql, []int{10000, 100000}, 1, 2)

	sql = "select * from test2 where (id >= ? and id <= ?) or id < ?"
	checkSharding(t, sql, []int{10000, 100000, 100}, 0, 1, 2)

	sql = "select * from test2 where (id >= ? and id <= ?) or (id < ? and name > ?)"
	checkSharding(t, sql, []int{10000, 100000, 100, 100000}, 0, 1, 2)

	sql = "select * from test2 where id in (?, ?)"
	checkSharding(t, sql, []int{1, 10000}, 0, 1)

	sql = "select * from test2 where id not in (?, ?)"
	checkSharding(t, sql, []int{1, 10000}, 0, 1, 2)

	sql = "select * from test2 where id in (?, ?)"
	checkSharding(t, sql, []int{1000, 10000}, 0, 1)

	sql = "select * from test2 where id > ?"
	checkSharding(t, sql, []int{-1}, 0, 1, 2)

	sql = "select * from test2 where id > ? and id < ?"
	checkSharding(t, sql, []int{-1, 11000}, 0, 1)
}

func TestValueSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (5)"
	checkSharding(t, sql, nil, 5)

	sql = "insert into test2 (id) values (10000)"
	checkSharding(t, sql, nil, 1)

	sql = "insert into test2 (id) values (20000)"
	checkSharding(t, sql, nil, 2)

	sql = "insert into test2 (id) values (200000)"
	checkSharding(t, sql, nil, 2)
}

func TestValueVarArgSharding(t *testing.T) {
	var sql string

	sql = "insert into test1 (id) values (?)"
	checkSharding(t, sql, []int{5}, 5)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{10000}, 1)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{20000}, 2)

	sql = "insert into test2 (id) values (?)"
	checkSharding(t, sql, []int{200000}, 2)
}
func TestBadUpdateExpr(t *testing.T) {
	var sql string

	r := newTestDBRule()

	sql = "insert into test1 (id) values (5) on duplicate key update  id = 10"

	if _, err := GetShardList(sql, r, nil); err == nil {
		t.Fatal("must err")
	}

	sql = "update test1 set id = 10 where id = 5"

	if _, err := GetShardList(sql, r, nil); err == nil {
		t.Fatal("must err")
	}
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
