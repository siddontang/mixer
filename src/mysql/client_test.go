package mysql

import (
	"testing"
)

var TestDBAddr = "10.20.188.113:3306"
var TestDBUser = "qing"
var TestDBPassword = "admin"
var TestDBName = "test"

func TestClient_Connect(t *testing.T) {
	client := NewClient()

	err := client.Connect(TestDBAddr, TestDBUser, TestDBPassword, TestDBName)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("capability: ", client.capability)
	t.Log("status: ", client.status)
	t.Log("authData: ", string(client.authData))
	t.Log("authName: ", client.authName)

	client.Close()
}
