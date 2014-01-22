package proxy

import (
	"fmt"
)

type Schema struct {
	cfg      *Config
	name     string
	nodes    []*DataNode
	rw_split bool
}

func NewSchema(cfg *Config, cfgSchema *ConfigSchema, nodes []*DataNode) *Schema {
	s := new(Schema)

	s.cfg = cfg
	s.name = cfgSchema.Name
	s.nodes = nodes
	s.rw_split = cfgSchema.RWSplit

	return s
}

func (s *Schema) SelectNode(sql string) *DataNode {
	return nil
}

type Schemas map[string]*Schema

func (ss Schemas) GetSchema(name string) *Schema {
	if s, ok := ss[name]; ok {
		return s
	} else {
		return nil
	}
}

func NewSchemas(cfg *Config, nodes DataNodes) Schemas {
	s := make(Schemas, len(cfg.Schemas))

	for _, v := range cfg.Schemas {
		if len(v.Nodes) == 0 {
			panic(fmt.Sprintf("schema %s has no node", v.Name))
		}

		nds := make([]*DataNode, 0, len(v.Nodes))
		for _, nodeName := range v.Nodes {
			if node := nodes.GetNode(nodeName); node == nil {
				panic(fmt.Sprintf("schema %s has invalid node name %s", v.Name, nodeName))
			} else {
				nds = append(nds, node)
			}
		}

		s[v.Name] = NewSchema(cfg, &v, nds)
	}

	return s
}
