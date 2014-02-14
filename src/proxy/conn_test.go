package proxy

import (
	"testing"
)

func TestConn_Handshake(t *testing.T) {
	db := newTestDB()

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
