package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	var testConfigData = []byte(
		`
addr : 127.0.0.1:4000
user : root
password : 
nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    rw_split: true
    user: root
    password:
    master : 127.0.0.1:3306
    master_backup : 127.0.0.1:3307
    slave : 127.0.0.1:4306
-
    name : node2 
    user: root
    password:
    master : 127.0.0.1:3308

schemas :
-
    db : mixer 
    nodes: [node1, node2]

rules:
-
    db: mixer
    table: test1 
    key: id
    type: hash
    nodes: node(1-2)
-
    db: mixer
    table: test2 
    key: name
    nodes: node1,node2
    type: range
    range: -6FFFFFFFFFFFFFFF-AFFFFFFFFFFFFFFF-
-   db: mixer
    table: 
    key:
    nodes: node1
    type: default
`)

	cfg, err := ParseConfigData(testConfigData)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Nodes) != 2 {
		t.Fatal(len(cfg.Nodes))
	}

	if len(cfg.Schemas) != 1 {
		t.Fatal(len(cfg.Schemas))
	}

	if len(cfg.Rules) != 3 {
		t.Fatal(len(cfg.Rules))
	}

	testNode := NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 300,
		IdleConns:        16,
		RWSplit:          true,

		User:     "root",
		Password: "",

		Master:       "127.0.0.1:3306",
		MasterBackup: "127.0.0.1:3307",
		Slave:        "127.0.0.1:4306",
	}

	if !reflect.DeepEqual(cfg.Nodes[0], testNode) {
		fmt.Printf("%v\n", cfg.Nodes[0])
		t.Fatal("node1 must equal")
	}

	testNode = NodeConfig{
		Name:   "node2",
		User:   "root",
		Master: "127.0.0.1:3308",
	}

	if !reflect.DeepEqual(cfg.Nodes[1], testNode) {
		t.Fatal("node2 must equal")
	}

	testSchema := SchemaConfig{
		DB:    "mixer",
		Nodes: []string{"node1", "node2"},
	}

	if !reflect.DeepEqual(cfg.Schemas[0], testSchema) {
		t.Fatal("schema must equal")
	}

	testRule := RuleConfig{
		DB:    "mixer",
		Table: "test1",
		Key:   "id",
		Nodes: "node(1-2)",
		Type:  "hash",
		Range: "",
	}

	if !reflect.DeepEqual(cfg.Rules[0], testRule) {
		t.Fatal("rule0 must equal")
	}

	testRule = RuleConfig{
		DB:    "mixer",
		Table: "test2",
		Key:   "name",
		Nodes: "node1,node2",
		Type:  "range",
		Range: "-6FFFFFFFFFFFFFFF-AFFFFFFFFFFFFFFF-",
	}

	if !reflect.DeepEqual(cfg.Rules[1], testRule) {
		t.Fatal("rule1 must equal")
	}

	testRule = RuleConfig{
		DB:    "mixer",
		Table: "",
		Key:   "",
		Nodes: "node1",
		Type:  "default",
		Range: "",
	}

	if !reflect.DeepEqual(cfg.Rules[2], testRule) {
		t.Fatal("rule2 must equal")
	}

	ruleConfig := cfg.NewRouterConfig()
	if len(ruleConfig.Rules) == 0 {
		t.Fatal("must not 0")
	}

	if ruleConfig.Rules[2].Type != "default" {
		t.Fatal(ruleConfig.Rules[2].Type)
	}
}
