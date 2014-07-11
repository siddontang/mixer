package server

import (
	"fmt"
	. "github.com/siddontang/mixer/parser"
)

type schema struct {
	server *Server
	cfg    *config
	db     string
	nodes  []*node

	masterNodes []*node
	slaveNodes  []*node
}

func newSchema(server *Server, cfgSchema *configSchema, nodes []*node) *schema {
	s := new(schema)

	s.server = server
	s.cfg = server.cfg
	s.db = cfgSchema.DB
	s.nodes = nodes

	s.masterNodes = make([]*node, 0)
	s.slaveNodes = make([]*node, 0)

	for _, n := range nodes {
		if n.Mode == MASTER_MODE {
			s.masterNodes = append(s.masterNodes, n)
		} else {
			s.slaveNodes = append(s.slaveNodes, n)
		}
	}

	return s
}

type routeQuery struct {
	Query string
	Args  []interface{}
}

//return a map key is node and value is the routeQuery the node will run
func (s *schema) Route(l *lex, inTrans bool) (map[*node]routeQuery, error) {
	//todo
	//rebuild query for different node

	//if in transaction, we will send all query to master nodes
	if inTrans {
		return s.routeNodes(l, s.masterNodes)
	}

	switch l.Get(0).Type {
	case TK_SQL_SELECT:
		hasFrom := false
		for _, t := range l.Tokens {
			if t.Type == TK_SQL_FROM {
				hasFrom = true
				break
			}
		}

		//if select have no from, we will route to master 0 or slave 0
		if !hasFrom {
			if len(s.masterNodes) > 0 {
				return map[*node]routeQuery{
					s.masterNodes[0]: routeQuery{l.Query, l.Args},
				}, nil
			} else if len(s.slaveNodes) > 0 {
				return map[*node]routeQuery{
					s.slaveNodes[0]: routeQuery{l.Query, l.Args},
				}, nil
			} else {
				return nil, fmt.Errorf("no proper node to route")
			}
		}

		//select may redirect to slave if slave exists
		if len(s.slaveNodes) > 0 {
			return s.routeNodes(l, s.slaveNodes)
		} else {
			return s.routeNodes(l, s.masterNodes)
		}
	default:
		return s.routeNodes(l, s.masterNodes)
	}
}

func (s *schema) routeNodes(l *lex, ns []*node) (map[*node]routeQuery, error) {
	if len(ns) == 0 {
		return nil, fmt.Errorf("no proper node to route")
	}

	//now we only return first node
	return map[*node]routeQuery{ns[0]: routeQuery{l.Query, l.Args}}, nil
}

type schemas map[string]*schema

func (ss schemas) GetSchema(db string) *schema {
	if s, ok := ss[db]; ok {
		return s
	} else {
		return nil
	}
}

func newSchemas(server *Server, nodes nodes) schemas {
	cfg := server.cfg

	s := make(schemas, len(cfg.Schemas))

	for _, v := range cfg.Schemas {
		if len(v.Nodes) == 0 {
			panic(fmt.Sprintf("schema %s has no node", v.DB))
		}

		nds := make([]*node, 0, len(v.Nodes))
		for _, nodeName := range v.Nodes {
			if node := nodes.GetNode(nodeName); node == nil {
				panic(fmt.Sprintf("schema %s has invalid node name %s", v.DB, nodeName))
			} else {
				nds = append(nds, node)
			}
		}

		s[v.DB] = newSchema(server, &v, nds)
	}

	return s
}
