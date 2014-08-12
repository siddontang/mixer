// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"strconv"

	"github.com/siddontang/mixer/router"
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
}

/*
	Limitation:

	where, eg, key name is id, only supports below now:

		where id = 1
		where id in (1, 2, 3)
		where id > 1
		where id >= 1
		where id < 1
		where id <= 1
		where id between 1 and 10
*/
func GetShardList(sql string, bindVariables map[string]interface{}, r *router.DBRules) (nodes []string, err error) {
	var stmt Statement
	stmt, err = Parse(sql)
	if err != nil {
		return nil, err
	}

	return GetStmtShardList(stmt, bindVariables, r)
}

func GetStmtShardList(stmt Statement, bindVariables map[string]interface{}, r *router.DBRules) (nodes []string, err error) {
	defer handleError(&err)

	plan := getRoutingPlan(stmt, r)

	ns := plan.shardListFromPlan(bindVariables)

	nodes = make([]string, 0, len(ns))
	for _, i := range ns {
		nodes = append(nodes, plan.rule.Nodes[i])
	}

	return nodes, nil
}

func (plan *RoutingPlan) shardListFromPlan(bindVariables map[string]interface{}) (shardList []int) {
	if plan.criteria == nil {
		return makeList(0, len(plan.rule.Nodes))
	}

	switch criteria := plan.criteria.(type) {
	case Values:
		index := plan.findInsertShard(criteria, bindVariables)
		return []int{index}
	case *ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>":
			index := plan.findShard(criteria.Right, bindVariables)
			return []int{index}
		case "<", "<=":
			if plan.rule.Type == router.HashRuleType {
				return makeList(0, len(plan.rule.Nodes))
			}

			index := plan.findShard(criteria.Right, bindVariables)
			return makeList(0, index+1)
		case ">", ">=":
			if plan.rule.Type == router.HashRuleType {
				return makeList(0, len(plan.rule.Nodes))
			}

			index := plan.findShard(criteria.Right, bindVariables)
			return makeList(index, len(plan.rule.Nodes))
		case "in":
			return plan.findShardList(criteria.Right, bindVariables)
		}
	case *RangeCond:
		if plan.rule.Type == router.HashRuleType {
			return makeList(0, len(plan.rule.Nodes))
		}

		if criteria.Operator == "between" {
			start := plan.findShard(criteria.From, bindVariables)
			last := plan.findShard(criteria.To, bindVariables)
			if last < start {
				start, last = last, start
			}
			return makeList(start, last+1)
		}
	}
	return makeList(0, len(plan.rule.Nodes))
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
		plan.criteria = plan.routingAnalyzeValues(stmt.Rows.(Values))
		return plan
	case *Replace:
		if _, ok := stmt.Rows.(SelectStatement); ok {
			panic(NewParserError("select in replace not allowed"))
		}

		plan.rule = r.GetRule(String(stmt.Table))
		plan.criteria = plan.routingAnalyzeValues(stmt.Rows.(Values))
		return plan

	case *Select:
		plan.rule = r.GetRule(String(stmt.From[0]))
		where = stmt.Where
	case *Update:
		plan.rule = r.GetRule(String(stmt.Table))
		where = stmt.Where
	case *Delete:
		plan.rule = r.GetRule(String(stmt.Table))
		where = stmt.Where
	}

	if where != nil {
		plan.criteria = plan.routingAnalyzeBoolean(where.Expr)
	} else {
		plan.rule = r.DefaultRule
	}
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

func (plan *RoutingPlan) routingAnalyzeBoolean(node BoolExpr) BoolExpr {
	switch node := node.(type) {
	case *AndExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		if left != nil && right != nil {
			return nil
		} else if left != nil {
			return left
		} else {
			return right
		}
	case *ParenBoolExpr:
		return plan.routingAnalyzeBoolean(node.Expr)
	case *ComparisonExpr:
		switch {
		case StringIn(node.Operator, "=", "<", ">", "<=", ">=", "<=>"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return node
			}
		case node.Operator == "in":
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				return node
			}
		}
	case *RangeCond:
		if node.Operator != "between" {
			return nil
		}
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return node
		}
	}
	return nil
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

func (plan *RoutingPlan) findShardList(valExpr ValExpr, bindVariables map[string]interface{}) []int {
	shardset := make(map[int]bool)
	switch node := valExpr.(type) {
	case ValTuple:
		for _, n := range node {
			index := plan.findShard(n, bindVariables)
			shardset[index] = true
		}
	}
	shardlist := make([]int, len(shardset))
	index := 0
	for k := range shardset {
		shardlist[index] = k
		index++
	}

	return shardlist
}

func (plan *RoutingPlan) findInsertShard(vals Values, bindVariables map[string]interface{}) int {
	index := -1
	for i := 0; i < len(vals); i++ {
		first_value_expression := vals[i].(ValTuple)[0]
		newIndex := plan.findShard(first_value_expression, bindVariables)
		if index == -1 {
			index = newIndex
		} else if index != newIndex {
			panic(NewParserError("insert has multiple shard targets"))
		}
	}
	return index
}

func (plan *RoutingPlan) findShard(valExpr ValExpr, bindVariables map[string]interface{}) int {
	value := getBoundValue(valExpr, bindVariables)
	return plan.rule.FindNodeIndex(value)
}

func getBoundValue(valExpr ValExpr, bindVariables map[string]interface{}) interface{} {
	switch node := valExpr.(type) {
	case ValTuple:
		if len(node) != 1 {
			panic(NewParserError("tuples not allowed as insert values"))
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return getBoundValue(node[0], bindVariables)
	case StrVal:
		return string(node)
	case NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			panic(NewParserError("%s", err.Error()))
		}
		return val
	case ValArg:
		value := findBindValue(node, bindVariables)
		return value
	}
	panic("Unexpected token")
}

func findBindValue(valArg ValArg, bindVariables map[string]interface{}) interface{} {
	if bindVariables == nil {
		panic(NewParserError("No bind variable for " + string(valArg)))
	}
	value, ok := bindVariables[string(valArg[1:])]
	if !ok {
		panic(NewParserError("No bind variable for " + string(valArg)))
	}
	return value
}

func makeList(start, end int) []int {
	list := make([]int, end-start)
	for i := start; i < end; i++ {
		list[i-start] = i
	}
	return list
}
