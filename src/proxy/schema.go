package proxy

import (
	"fmt"
)

type schema struct {
	server   *Server
	cfg      *config
	db       string
	nodes    []*node
	rw_split bool
}

func newSchema(server *Server, cfgSchema *configSchema, nodes []*node) *schema {
	s := new(schema)

	s.server = server
	s.cfg = server.cfg
	s.db = cfgSchema.DB
	s.nodes = nodes
	s.rw_split = cfgSchema.RWSplit

	return s
}

//return a map key is datanode and value is the query the datanode will run
func (s *schema) Route(query []byte) (map[*node][]byte, error) {
	//todo
	//1, parse query with rule
	//2, rebuild query to send different datanode

	//now we only return first datanode

	return map[*node][]byte{s.nodes[0]: query}, nil
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
