package proxy

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c, err := NewConfig("../etc/proxy.json")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c.Address)
}
