package proxy

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	_, err := newConfig("../../etc/proxy")
	if err != nil {
		t.Fatal(err)
	}
}
