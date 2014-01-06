package mysql

import (
	"testing"
)

func TestClient_InitDB(t *testing.T) {
	client := newTestClient()

	err := client.UseDB("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Ping(t *testing.T) {
	client := newTestClient()

	err := client.Ping()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Quit(t *testing.T) {
	client := newTestClient()

	client.Quit()

	err := client.ReConnect()
	if err != nil {
		t.Fatal(err)
	}
}
