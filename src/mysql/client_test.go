package mysql

import (
	"github.com/siddontang/golib/log"
	"sync"
	"testing"
)

var TestDBAddr = "10.20.188.113:3306"
var TestDBUser = "qing"
var TestDBPassword = "admin"
var TestDBName = "test"

var testClient *Client

var testClientOnce sync.Once

func newTestClient() *Client {
	f := func() {
		testClient = NewClient()

		err := testClient.Connect(TestDBAddr, TestDBUser, TestDBPassword, TestDBName)

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	testClientOnce.Do(f)

	return testClient
}

func TestClient_Connect(t *testing.T) {
	client := newTestClient()

	t.Log("capability: ", client.capability)
	t.Log("status: ", client.status)
	t.Log("authData: ", string(client.authData))
	t.Log("authName: ", client.authName)
}
