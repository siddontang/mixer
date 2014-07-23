// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"strconv"

	"github.com/siddontang/mixer/router"
)

const (
	ROUTE_BY_CONDITION = iota
	ROUTE_BY_VALUE
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

type RoutingPlan struct {
	rule        *router.Rule
	routingType int
	criteria    *Node
}

func GetShardList(sql string, bindVariables map[string]interface{}, r *router.DBRules) (nodes []string, err error) {
	defer handleError(&err)

	plan := buildPlan(sql, r)
	if plan.rule.Type == router.DefaultRuleType {
		return plan.rule.Nodes, nil
	}

	ns := shardListFromPlan(plan, bindVariables)

	nodes = make([]string, 0, len(ns))
	for _, i := range ns {
		nodes = append(nodes, plan.rule.Nodes[i])
	}

	return nodes, nil
}

func buildPlan(sql string, r *router.DBRules) (plan *RoutingPlan) {
	statement, err := Parse(sql)
	if err != nil {
		panic(err)
	}
	return getRoutingPlan(statement, r)
}

func shardListFromPlan(plan *RoutingPlan, bindVariables map[string]interface{}) (shardList []int) {
	r := plan.rule

	if plan.routingType == ROUTE_BY_VALUE {
		index := plan.criteria.findInsertShard(bindVariables, r)
		return []int{index}
	}

	if plan.criteria == nil {
		return makeList(0, len(r.Nodes))
	}

	switch plan.criteria.Type {
	case '=', NULL_SAFE_EQUAL:
		index := plan.criteria.NodeAt(1).findShard(bindVariables, r)
		return []int{index}
	case '<', LE:
		if r.Type == router.HashRuleType {
			return makeList(0, len(r.Nodes))
		}

		index := plan.criteria.NodeAt(1).findShard(bindVariables, r)
		return makeList(0, index+1)
	case '>', GE:
		if r.Type == router.HashRuleType {
			return makeList(0, len(r.Nodes))
		}

		index := plan.criteria.NodeAt(1).findShard(bindVariables, r)
		return makeList(index, len(r.Nodes))
	case IN:
		return plan.criteria.NodeAt(1).findShardList(bindVariables, r)
	case BETWEEN:
		if r.Type == router.HashRuleType {
			return makeList(0, len(r.Nodes))
		}

		start := plan.criteria.NodeAt(1).findShard(bindVariables, r)
		last := plan.criteria.NodeAt(2).findShard(bindVariables, r)
		if last < start {
			start, last = last, start
		}
		return makeList(start, last+1)
	}
	return makeList(0, len(r.Nodes))
}

func getRoutingPlan(statement Statement, r *router.DBRules) (plan *RoutingPlan) {
	plan = &RoutingPlan{}
	var tableNode *Node
	if ins, ok := statement.(*Insert); ok {
		tableNode = ins.Table
		if _, ok := ins.Values.(SelectStatement); ok {
			panic(NewParserError("select in insert not allowed"))
		}

		plan.rule = r.GetRule(tableNode.String())

		plan.routingType = ROUTE_BY_VALUE
		plan.criteria = ins.Values.(*Node).NodeAt(0).routingAnalyzeValues(plan.rule)
		return plan
	}
	var where *Node
	plan.routingType = ROUTE_BY_CONDITION
	switch stmt := statement.(type) {
	case *Select:
		//now only support from only one table
		tableNode = stmt.From[0]
		where = stmt.Where
	case *Update:
		tableNode = stmt.Table
		where = stmt.Where
	case *Delete:
		tableNode = stmt.Table
		where = stmt.Where
	}

	plan.rule = r.GetRule(tableNode.String())

	if where != nil && where.Len() > 0 {
		plan.criteria = where.NodeAt(0).routingAnalyzeBoolean(plan.rule)
	}

	return plan
}

func (node *Node) routingAnalyzeValues(r *router.Rule) *Node {
	// Analyze first value of every item in the list
	for i := 0; i < node.Len(); i++ {
		value_expression_list := node.NodeAt(i)
		inner_list, ok := value_expression_list.At(0).(*Node)
		if !ok {
			panic(NewParserError("insert is too complex"))
		}
		result := inner_list.NodeAt(0).routingAnalyzeValue(r)
		if result != VALUE_NODE {
			panic(NewParserError("insert is too complex"))
		}
	}
	return node
}

func (node *Node) routingAnalyzeBoolean(r *router.Rule) *Node {
	switch node.Type {
	case AND:
		left := node.NodeAt(0).routingAnalyzeBoolean(r)
		right := node.NodeAt(1).routingAnalyzeBoolean(r)
		if left != nil && right != nil {
			return nil
		} else if left != nil {
			return left
		} else {
			return right
		}
	case '(':
		sub, ok := node.At(0).(*Node)
		if !ok {
			return nil
		}
		return sub.routingAnalyzeBoolean(r)
	case '=', '<', '>', LE, GE, NULL_SAFE_EQUAL:
		left := node.NodeAt(0).routingAnalyzeValue(r)
		right := node.NodeAt(1).routingAnalyzeValue(r)
		if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
			return node
		}
	case IN:
		left := node.NodeAt(0).routingAnalyzeValue(r)
		right := node.NodeAt(1).routingAnalyzeValue(r)
		if left == EID_NODE && right == LIST_NODE {
			return node
		}
	case BETWEEN:
		left := node.NodeAt(0).routingAnalyzeValue(r)
		right1 := node.NodeAt(1).routingAnalyzeValue(r)
		right2 := node.NodeAt(2).routingAnalyzeValue(r)
		if left == EID_NODE && right1 == VALUE_NODE && right2 == VALUE_NODE {
			return node
		}
	}
	return nil
}

func (node *Node) routingAnalyzeValue(r *router.Rule) int {
	switch node.Type {
	case ID:
		if string(node.Value) == r.Key {
			return EID_NODE
		}
	case '.':
		return node.NodeAt(1).routingAnalyzeValue(r)
	case '(':
		sub, ok := node.At(0).(*Node)
		if !ok {
			return OTHER_NODE
		}
		return sub.routingAnalyzeValue(r)
	case NODE_LIST:
		for i := 0; i < node.Len(); i++ {
			if node.NodeAt(i).routingAnalyzeValue(r) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE
	case STRING, NUMBER, VALUE_ARG:
		return VALUE_NODE
	}
	return OTHER_NODE
}

func (node *Node) findShardList(bindVariables map[string]interface{}, r *router.Rule) []int {
	shardset := make(map[int]bool)
	switch node.Type {
	case '(':
		return node.NodeAt(0).findShardList(bindVariables, r)
	case NODE_LIST:
		for i := 0; i < node.Len(); i++ {
			index := node.NodeAt(i).findShard(bindVariables, r)
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

func (node *Node) findInsertShard(bindVariables map[string]interface{}, r *router.Rule) int {
	index := -1
	for i := 0; i < node.Len(); i++ {
		first_value_expression := node.NodeAt(i).NodeAt(0).NodeAt(0) // '('->value_expression_list->first_value
		newIndex := first_value_expression.findShard(bindVariables, r)
		if index == -1 {
			index = newIndex
		} else if index != newIndex {
			panic(NewParserError("insert has multiple shard targets"))
		}
	}
	return index
}

func (node *Node) findShard(bindVariables map[string]interface{}, r *router.Rule) int {
	value := node.getBoundValue(bindVariables)
	return r.FindNodeIndex(value)
}

func (node *Node) getBoundValue(bindVariables map[string]interface{}) interface{} {
	switch node.Type {
	case '(':
		return node.NodeAt(0).getBoundValue(bindVariables)
	case STRING:
		return string(node.Value)
	case NUMBER:
		val, err := strconv.ParseInt(string(node.Value), 10, 64)
		if err != nil {
			panic(NewParserError("%s", err.Error()))
		}
		return val
	case VALUE_ARG:
		value := node.findBindValue(bindVariables)
		return value
	}
	panic("Unexpected token")
}

func (node *Node) findBindValue(bindVariables map[string]interface{}) interface{} {
	if bindVariables == nil {
		panic(NewParserError("No bind variable for " + string(node.Value)))
	}
	value, ok := bindVariables[string(node.Value[1:])]
	if !ok {
		panic(NewParserError("No bind variable for " + string(node.Value)))
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
