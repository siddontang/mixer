package router

import (
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type Rule struct {
	DB    string
	Table string
	Key   string

	Type string

	Nodes []string
	Shard Shard
}

func (r *Rule) FindNode(key interface{}) string {
	i := r.Shard.FindForKey(key)
	return r.Nodes[i]
}

func (r *Rule) FindNodeIndex(key interface{}) int {
	return r.Shard.FindForKey(key)
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s:%s", r.DB, r.Table)
}

type DBRules struct {
	DB          string
	Rules       map[string]*Rule
	DefaultRule *Rule
}

func NewDefaultDBRules(db string, node string) *DBRules {
	r := new(DBRules)
	r.DB = db
	r.Rules = make(map[string]*Rule)
	r.DefaultRule = &Rule{
		DB:    db,
		Type:  DefaultRuleType,
		Nodes: []string{node}}
	return r
}

func (r *DBRules) GetRule(table string) *Rule {
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
}

func (r *DBRules) String() string {
	return r.DB
}

type Router struct {
	Rules map[string]*DBRules
}

func (r *Router) GetDBRules(db string) *DBRules {
	return r.Rules[db]
}

func (r *Router) GetRule(db string, table string) *Rule {
	dbRule := r.Rules[db]
	if dbRule == nil {
		return nil
	}

	return dbRule.GetRule(table)
}

func NewRouter(cfg *Config) (*Router, error) {
	rt := new(Router)
	rt.Rules = make(map[string]*DBRules, len(cfg.Rules))
	for _, r := range cfg.Rules {
		rule, err := r.ParseRule()
		if err != nil {
			return nil, err
		}

		dbRules, ok := rt.Rules[rule.DB]
		if !ok {
			dbRules = &DBRules{DB: rule.DB}
			dbRules.Rules = make(map[string]*Rule)
			rt.Rules[rule.DB] = dbRules
		}

		if rule.Type == DefaultRuleType {
			if dbRules.DefaultRule != nil {
				return nil, fmt.Errorf("default rule duplicate, must only one")
			}
			dbRules.DefaultRule = rule
		} else {
			if _, ok := dbRules.Rules[rule.Table]; ok {
				return nil, fmt.Errorf("table %s rule in %s duplicate", rule.Table, rule.DB)
			}
			dbRules.Rules[rule.Table] = rule
		}
	}
	return rt, nil
}

func NewRouterConfigData(data []byte) (*Router, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	return NewRouter(&cfg)
}

func NewRouterConfigFile(fileName string) (*Router, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	return NewRouter(&cfg)
}
