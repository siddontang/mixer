package proxy

import (
	"encoding/json"
	"github.com/siddontang/golib/log"
	"io/ioutil"
)

type Config struct {
	Address  string   `json:"address"`
	Masters  []string `json:"backend-addresses"`
	Salves   []string `json:"read-only-backend-addresses"`
	User     string   `json:"user"`
	Password string   `json:"password"`
}

func NewConfig(configFile string) (*Config, error) {
	log.Info("new config %s", configFile)

	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	var c Config

	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
