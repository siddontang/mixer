package server

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	_, err := newConfig("../../etc/proxy.json")
	if err != nil {
		t.Fatal(err)
	}
}
