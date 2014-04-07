package proxy

import (
	"encoding/json"
	"io/ioutil"
)

type configDataNode struct {
	Name               string   `json:"name"`
	DSN                []string `json:dsn`
	Mode               string   `json:"mode"`
	SwitchAfterNoAlive int      `json:"switch_after_noalive"`
	MaxIdleConns       int      `json:"max_idle_conns"`
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
	c := new(config)

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
