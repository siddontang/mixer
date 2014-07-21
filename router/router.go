package router

import (
	"fmt"
)

type Rule struct {
	DB    string
	Table string
	Key   string

	Type string

	Nodes []string
	Shard Shard
}

func (r *Rule) FindForKey(key interface{}) string {
	i := r.Shard.FindForKey(key)
	return r.Nodes[i]
}

type DBRules struct {
	DB          string
	Rules       map[string]*Rule
	DefaultRule *Rule
}

func (r *DBRules) GetRule(table string) *Rule {
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
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
