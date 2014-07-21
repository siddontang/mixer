package router

import (
	"regexp"
	"strings"
)

var (
	HashShardType  = "hash"
	RangeShardType = "range"
)

type RuleConfig struct {
	DB    string `json:"db"`
	Table string `json:"table"`
	Key   string `json:"key"`
	Nodes string `json:"nodes"`
	Range string `json:"range"`
}

type Rule struct {
	DB    string
	Table string
	Key   string

	Nodes []Node
	Shard Shard
}

func (c *RuleConfig) ParseRule() (*Rule, error) {
	r := new(Rule)
	r.DB = c.DB
	r.Table = c.Table
	r.Key = c.Key

	if err := c.parseNodes(); err != nil {
		return nil, err
	}

	if err := c.parseShard(); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *RuleConfig) parseNodes() error {
	r, err := regexp.Compile("(%d-%d)")
	if err != nil {
		return err
	}

	ns := strings.Split(c.Nodes, ",")

	var nodes map[string]struct{}
	for _, n := range ns {
		n = strings.TrimSpace(n)
		if s := r.FindAllString(n, -1); s == nil {

		}

	}

	return nil
}

func (c *RuleConfig) parseShard() error {
	return nil
}
