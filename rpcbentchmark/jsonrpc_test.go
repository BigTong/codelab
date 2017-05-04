package rpcbm

import (
	"testing"
)

func TestJsonRpcUserUpdate(t *testing.T) {
	StartServer()

	user := NewFakeUser()
	client := NewJsonRpcClient(JSONRPC_SERVER_ADDRESS, 5)
	if client == nil {
		t.Error("get nil client")
	}
	defer client.Close()

	respUser, err := client.UpdateUser(user)

	if err != nil {
		t.Error(err.Error())
	}

	if respUser.Name != USER_MOCK_NAME {
		t.Error("not update user name")
	}

}
