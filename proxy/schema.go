package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/router"
)

type Schema struct {
	db string

	nodes map[string]*Node

	rule *router.DBRules
}

func (s *Server) parseSchemas() error {
	r, err := router.NewRouter(s.cfg.NewRouterConfig())
	if err != nil {
		return err
	}

	for _, v := range s.cfg.Schemas {
		if _, ok := s.schemas[v.DB]; ok {
			return fmt.Errorf("duplicate schema %s", v.DB)
		}

		if len(v.Nodes) == 0 {
			return fmt.Errorf("schema %s must have a node", v.DB)
		}

		nodes := make(map[string]*Node)
		for _, n := range v.Nodes {
			if s.getNode(n) == nil {
				return fmt.Errorf("schema %s node %s is not exists", v.DB, n)
			}

			if _, ok := nodes[n]; ok {
				return fmt.Errorf("schema %s node %s duplicate", v.DB, n)
			}

			nodes[n] = s.getNode(n)
		}

		dbRules := r.GetDBRules(v.DB)
		if dbRules == nil {
			if len(v.Nodes) != 1 {
				return fmt.Errorf("schema %s must be set a rule for multi nodes %v", v.DB, v.Nodes)
			} else {
				dbRules = router.NewDefaultDBRules(v.DB, v.Nodes[0])
			}
		}

		for _, v := range dbRules.Rules {
			for _, n := range v.Nodes {
				if s.getNode(n) == nil {
					return fmt.Errorf("rule %s node %s is not exists", v, n)
				}
			}
		}

		if s.getNode(dbRules.DefaultRule.Nodes[0]) == nil {
			return fmt.Errorf("rule %s node %s is not exists", v, dbRules.DefaultRule.Nodes[0])
		}

		s.schemas[v.DB] = &Schema{
			db:    v.DB,
			nodes: nodes,
			rule:  dbRules,
		}
	}

	return nil
}

func (s *Server) getSchema(db string) *Schema {
	return s.schemas[db]
}
