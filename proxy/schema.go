package proxy

import (
	"fmt"
	"github.com/siddontang/mixer/router"
)

type Schema struct {
	db string

	nodes map[string]*Node

	rule *router.Router
}

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	for _, schemaCfg := range s.cfg.Schemas {
		if _, ok := s.schemas[schemaCfg.DB]; ok {
			return fmt.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
		if len(schemaCfg.Nodes) == 0 {
			return fmt.Errorf("schema [%s] must have a node.", schemaCfg.DB)
		}

		nodes := make(map[string]*Node)
		for _, n := range schemaCfg.Nodes {
			if s.getNode(n) == nil {
				return fmt.Errorf("schema [%s] node [%s] config is not exists.", schemaCfg.DB, n)
			}

			if _, ok := nodes[n]; ok {
				return fmt.Errorf("schema [%s] node [%s] duplicate.", schemaCfg.DB, n)
			}

			nodes[n] = s.getNode(n)
		}

		rule, err := router.NewRouter(&schemaCfg)
		if err != nil {
			return err
		}

		s.schemas[schemaCfg.DB] = &Schema{
			db:    schemaCfg.DB,
			nodes: nodes,
			rule:  rule,
		}
	}

	return nil
}

func (s *Server) getSchema(db string) *Schema {
	return s.schemas[db]
}
