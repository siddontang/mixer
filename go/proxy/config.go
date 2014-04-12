package proxy

import (
	"encoding/json"
	"io/ioutil"
)

type configDataNode struct {
	Name               string `json:"name"`
	Mode               string `json:"mode"`
	SwitchAfterNoAlive int    `json:"switch_after_noalive"`

	Backends []json.RawMessage `json:"backends"`
}

type configSchema struct {
	DB    string   `json:"db"`
	Nodes []string `json:"nodes"`
}

type config struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`

	Nodes []configDataNode `json:"nodes"`

	Schemas []configSchema `json:"schemas"`
}

func newConfig(configFile string) (*config, error) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	return newConfigJson(b)
}

func newConfigJson(configJson json.RawMessage) (*config, error) {
	c := new(config)

	err := json.Unmarshal(configJson, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
