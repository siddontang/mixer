package proxy

import (
	"fmt"
)

type Schema struct {
	server   *Server
	cfg      *Config
	name     string
	nodes    []*DataNode
	rw_split bool
}

func NewSchema(server *Server, cfgSchema *ConfigSchema, nodes []*DataNode) *Schema {
	s := new(Schema)

	s.server = server
	s.cfg = server.cfg
	s.name = cfgSchema.Name
	s.nodes = nodes
	s.rw_split = cfgSchema.RWSplit

	return s
}

//return a map key is datanode and value is the query the datanode will run
func (s *Schema) Route(query []byte) (map[*DataNode][]byte, error) {
	//todo
	//1, parse query with rule
	//2, rebuild query to send different datanode

	//now we only return first datanode

	return map[*DataNode][]byte{s.nodes[0]: query}, nil
}

type Schemas map[string]*Schema

func (ss Schemas) GetSchema(name string) *Schema {
	if s, ok := ss[name]; ok {
		return s
	} else {
		return nil
	}
}

func NewSchemas(server *Server, nodes DataNodes) Schemas {
	cfg := server.cfg

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

		s[v.Name] = NewSchema(server, &v, nds)
	}

	return s
}
