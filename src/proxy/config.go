package proxy

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"path"
)

type ConfigDataNode struct {
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Mode     string `yaml:"mode"`
}

type ConfigSchema struct {
	Name    string   `yaml:"name"`
	Nodes   []string `yaml:"nodes"`
	RWSplit bool     `yaml:"rw_split"`
}

type ConfigServer struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`

	MaxIdleConns int `yaml:"max_idle_conns"`

	DataNodes []ConfigDataNode `yaml:"datanodes"`

	Schemas []ConfigSchema `yaml:"schemas"`
}

type Config struct {
	ConfigServer
}

func (c *Config) loadServer(configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = goyaml.Unmarshal(b, &c.ConfigServer)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) loadRule(configFile string) error {
	return nil
}

func NewConfig(configDir string) (*Config, error) {
	c := new(Config)

	if err := c.loadServer(path.Join(configDir, "server.yaml")); err != nil {
		return nil, err
	}

	if err := c.loadRule(path.Join(configDir, "rule.yaml")); err != nil {
		return nil, err
	}

	return c, nil
}
