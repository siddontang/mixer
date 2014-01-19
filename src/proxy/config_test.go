package proxy

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	_, err := NewConfig("../etc/proxy")
	if err != nil {
		t.Fatal(err)
	}
}
