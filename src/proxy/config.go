package proxy

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type configDataNode struct {
	Name         string `json:"name"`
	Addr         string `json:"addr"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DB           string `json:"db"`
	Mode         string `json:"mode"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type configSchema struct {
	DB      string   `json:"db"`
	Nodes   []string `json:"nodes"`
	RWSplit bool     `json:"rw_split"`
}

type configServer struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`

	Nodes []configDataNode `json:"nodes"`

	Schemas []configSchema `json:"schemas"`
}

type config struct {
	configServer
}

func (c *config) loadServer(configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &c.configServer)
	if err != nil {
		return err
	}

	return nil
}

func (c *config) loadRule(configFile string) error {
	return nil
}

func newConfig(configDir string) (*config, error) {
	c := new(config)

	if err := c.loadServer(path.Join(configDir, "server.json")); err != nil {
		return nil, err
	}

	if err := c.loadRule(path.Join(configDir, "rule.json")); err != nil {
		return nil, err
	}

	return c, nil
}
