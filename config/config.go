package config

import (
	"github.com/siddontang/go-yaml/yaml"
	"io/ioutil"
)

type NodeConfig struct {
	Name             string `yaml:"name"`
	DownAfterNoAlive int    `yaml:"down_after_noalive"`
	IdleConns        int    `yaml:"idle_conns"`
	RWSplit          bool   `yaml:"rw_split"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`

	Master       string `yaml:"master"`
	MasterBackup string `yaml:"master_backup"`
	Slave        string `yaml:"slave"`
}

type SchemaConfig struct {
	DB    string   `yaml:"db"`
	Nodes []string `yaml:"nodes"`
}

type RuleConfig struct {
	DB    string `yaml:"db"`
	Table string `yaml:"table"`
	Key   string `yaml:"key"`
	Nodes string `yaml:"nodes"`
	Type  string `yaml:"type"`
	Range string `yaml:"range"`
}

type Config struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`

	Nodes []NodeConfig `yaml:"nodes"`

	Schemas []SchemaConfig `yaml:"schemas"`

	Rules []RuleConfig `yaml:"rules"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
