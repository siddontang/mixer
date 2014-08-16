// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"github.com/siddontang/mixer/router"
	"sort"
	"strconv"
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

type RoutingPlan struct {
	rule *router.Rule

	criteria SQLNode

	fullList []int

	bindVars map[string]interface{}
}

/*
	Limitation:

	where, eg, key name is id:

		where id = 1
		where id in (1, 2, 3)
		where id > 1
		where id >= 1
		where id < 1
		where id <= 1
		where id between 1 and 10
		where id >= 1 and id < 10
*/
func GetShardList(sql string, r *router.DBRules, bindVars map[string]interface{}) (nodes []string, err error) {
	var stmt Statement
	stmt, err = Parse(sql)
	if err != nil {
		return nil, err
	}

	return GetStmtShardList(stmt, r, bindVars)
}

func GetShardListIndex(sql string, r *router.DBRules, bindVars map[string]interface{}) (nodes []int, err error) {
	var stmt Statement
	stmt, err = Parse(sql)
	if err != nil {
		return nil, err
	}

	return GetStmtShardListIndex(stmt, r, bindVars)
}

func GetStmtShardList(stmt Statement, r *router.DBRules, bindVars map[string]interface{}) (nodes []string, err error) {
	defer handleError(&err)

	plan := getRoutingPlan(stmt, r)

	plan.bindVars = bindVars

	ns := plan.shardListFromPlan()

	nodes = make([]string, 0, len(ns))
	for _, i := range ns {
		nodes = append(nodes, plan.rule.Nodes[i])
	}

	return nodes, nil
}

func GetStmtShardListIndex(stmt Statement, r *router.DBRules, bindVars map[string]interface{}) (nodes []int, err error) {
	defer handleError(&err)

	plan := getRoutingPlan(stmt, r)

	plan.bindVars = bindVars

	ns := plan.shardListFromPlan()

	return ns, nil
}

func (plan *RoutingPlan) notList(l []int) []int {
	return differentList(plan.fullList, l)
}

func (plan *RoutingPlan) findConditionShard(expr BoolExpr) (shardList []int) {
	var index int
	switch criteria := expr.(type) {
	case *ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>":
			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
			} else {
				index = plan.findShard(criteria.Left)
			}
			return []int{index}
		case "<", "<=":
			if plan.rule.Type == router.HashRuleType {
				return plan.fullList
			}

			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
				if criteria.Operator == "<" {
					index = plan.adjustShardIndex(criteria.Right, index)
				}

				return makeList(0, index+1)
			} else {
				index = plan.findShard(criteria.Left)
				return makeList(index, len(plan.rule.Nodes))
			}
		case ">", ">=":
			if plan.rule.Type == router.HashRuleType {
				return plan.fullList
			}

			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
				return makeList(index, len(plan.rule.Nodes))
			} else {
				index = plan.findShard(criteria.Left)

				if criteria.Operator == ">" {
					index = plan.adjustShardIndex(criteria.Left, index)
				}
				return makeList(0, index+1)
			}
		case "in":
			return plan.findShardList(criteria.Right)
		case "not in":
			if plan.rule.Type == router.RangeRuleType {
				return plan.fullList
			}

			l := plan.findShardList(criteria.Right)
			return plan.notList(l)
		}
	case *RangeCond:
		if plan.rule.Type == router.HashRuleType {
			return plan.fullList
		}

		start := plan.findShard(criteria.From)
		last := plan.findShard(criteria.To)

		if criteria.Operator == "between" {
			if last < start {
				start, last = last, start
			}
			l := makeList(start, last+1)
			return l
		} else {
			if last < start {
				start, last = last, start
				start = plan.adjustShardIndex(criteria.To, start)
			} else {
				start = plan.adjustShardIndex(criteria.From, start)
			}

			l1 := makeList(0, start+1)
			l2 := makeList(last, len(plan.rule.Nodes))
			return unionList(l1, l2)
		}
	default:
		return plan.fullList
	}

	return plan.fullList
}

func (plan *RoutingPlan) shardListFromPlan() (shardList []int) {
	if plan.criteria == nil {
		return plan.fullList
	}

	//default rule will route all sql to one node
	//if rule has one node, we also can route directly
	if plan.rule.Type == router.DefaultRuleType || len(plan.rule.Nodes) == 1 {
		if len(plan.fullList) != 1 {
			panic(NewParserError("invalid rule nodes num %d, must 1", plan.fullList))
		}
		return plan.fullList
	}

	switch criteria := plan.criteria.(type) {
	case Values:
		index := plan.findInsertShard(criteria)
		return []int{index}
	case BoolExpr:
		return plan.routingAnalyzeBoolean(criteria)
	default:
		return plan.fullList
	}
}

func checkUpdateExprs(exprs UpdateExprs, rule *router.Rule) {
	if rule.Type == router.DefaultRuleType {
		return
	} else if len(rule.Nodes) == 1 {
		return
	}

	for _, e := range exprs {
		if string(e.Name.Name) == rule.Key {
			panic(NewParserError("routing key can not in update expression"))
		}
	}
}

func getRoutingPlan(statement Statement, r *router.DBRules) (plan *RoutingPlan) {
	plan = &RoutingPlan{}
	var where *Where
	switch stmt := statement.(type) {
	case *Insert:
		if _, ok := stmt.Rows.(SelectStatement); ok {
			panic(NewParserError("select in insert not allowed"))
		}

		plan.rule = r.GetRule(String(stmt.Table))

		if stmt.OnDup != nil {
			checkUpdateExprs(UpdateExprs(stmt.OnDup), plan.rule)
		}

		plan.criteria = plan.routingAnalyzeValues(stmt.Rows.(Values))
		plan.fullList = makeList(0, len(plan.rule.Nodes))
		return plan
	case *Replace:
		if _, ok := stmt.Rows.(SelectStatement); ok {
			panic(NewParserError("select in replace not allowed"))
		}

		plan.rule = r.GetRule(String(stmt.Table))
		plan.criteria = plan.routingAnalyzeValues(stmt.Rows.(Values))
		plan.fullList = makeList(0, len(plan.rule.Nodes))
		return plan

	case *Select:
		plan.rule = r.GetRule(String(stmt.From[0]))
		where = stmt.Where
	case *Update:
		plan.rule = r.GetRule(String(stmt.Table))

		checkUpdateExprs(stmt.Exprs, plan.rule)

		where = stmt.Where
	case *Delete:
		plan.rule = r.GetRule(String(stmt.Table))
		where = stmt.Where
	}

	if where != nil {
		plan.criteria = where.Expr
	} else {
		plan.rule = r.DefaultRule
	}
	plan.fullList = makeList(0, len(plan.rule.Nodes))

	return plan
}

func (plan *RoutingPlan) routingAnalyzeValues(vals Values) Values {
	// Analyze first value of every item in the list
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case ValTuple:
			result := plan.routingAnalyzeValue(tuple[0])
			if result != VALUE_NODE {
				panic(NewParserError("insert is too complex"))
			}
		default:
			panic(NewParserError("insert is too complex"))
		}
	}
	return vals
}

func (plan *RoutingPlan) routingAnalyzeBoolean(node BoolExpr) []int {
	switch node := node.(type) {
	case *AndExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)

		return interList(left, right)
	case *OrExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		return unionList(left, right)
	case *ParenBoolExpr:
		return plan.routingAnalyzeBoolean(node.Expr)
	case *ComparisonExpr:
		switch {
		case StringIn(node.Operator, "=", "<", ">", "<=", ">=", "<=>"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return plan.findConditionShard(node)
			}
		case StringIn(node.Operator, "in", "not in"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				return plan.findConditionShard(node)
			}
		}
	case *RangeCond:
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.findConditionShard(node)
		}
	}
	return plan.fullList
}

func (plan *RoutingPlan) routingAnalyzeValue(valExpr ValExpr) int {
	switch node := valExpr.(type) {
	case *ColName:
		if string(node.Name) == plan.rule.Key {
			return EID_NODE
		}
	case ValTuple:
		for _, n := range node {
			if plan.routingAnalyzeValue(n) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE
	case StrVal, NumVal, ValArg:
		return VALUE_NODE
	}
	return OTHER_NODE
}

func (plan *RoutingPlan) findShardList(valExpr ValExpr) []int {
	shardset := make(map[int]bool)
	switch node := valExpr.(type) {
	case ValTuple:
		for _, n := range node {
			index := plan.findShard(n)
			shardset[index] = true
		}
	}
	shardlist := make([]int, len(shardset))
	index := 0
	for k := range shardset {
		shardlist[index] = k
		index++
	}

	sort.Ints(shardlist)
	return shardlist
}

func (plan *RoutingPlan) findInsertShard(vals Values) int {
	index := -1
	for i := 0; i < len(vals); i++ {
		first_value_expression := vals[i].(ValTuple)[0]
		newIndex := plan.findShard(first_value_expression)
		if index == -1 {
			index = newIndex
		} else if index != newIndex {
			panic(NewParserError("insert has multiple shard targets"))
		}
	}
	return index
}

func (plan *RoutingPlan) findShard(valExpr ValExpr) int {
	value := plan.getBoundValue(valExpr)
	return plan.rule.FindNodeIndex(value)
}

func (plan *RoutingPlan) adjustShardIndex(valExpr ValExpr, index int) int {
	value := plan.getBoundValue(valExpr)

	s, ok := plan.rule.Shard.(router.RangeShard)
	if !ok {
		return index
	}

	if s.EqualStart(value, index) {
		index--
		if index < 0 {
			panic(NewParserError("invalid range sharding"))
		}
	}
	return index
}

func (plan *RoutingPlan) getBoundValue(valExpr ValExpr) interface{} {
	switch node := valExpr.(type) {
	case ValTuple:
		if len(node) != 1 {
			panic(NewParserError("tuples not allowed as insert values"))
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return plan.getBoundValue(node[0])
	case StrVal:
		return string(node)
	case NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			panic(NewParserError("%s", err.Error()))
		}
		return val
	case ValArg:
		return plan.bindVars[string(node[1:])]
	}
	panic("Unexpected token")
}

func makeList(start, end int) []int {
	list := make([]int, end-start)
	for i := start; i < end; i++ {
		list[i-start] = i
	}
	return list
}

// l1 & l2
func interList(l1 []int, l2 []int) []int {
	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}

	l3 := make([]int, 0, len(l1)+len(l2))
	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			l3 = append(l3, l1[i])
			i++
			j++
		} else if l1[i] < l2[j] {
			i++
		} else {
			j++
		}
	}

	return l3
}

// l1 | l2
func unionList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1)+len(l2))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			l3 = append(l3, l2[j])
			j++
		} else {
			l3 = append(l3, l1[i])
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	} else if j != len(l2) {
		l3 = append(l3, l2[j:]...)
	}

	return l3
}

// l1 - l2
func differentList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return []int{}
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			j++
		} else {
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	}

	return l3
}
