package rpcbm

import (
	"testing"
)

func TestGrpcUserUpdate(t *testing.T) {
	StartServer()

	user := NewFakeProtoUser()
	client := NewGrpcClient(GRPC_SERVER_ADDRESS)
	if client == nil {
		t.Error("get nil client")
	}
	defer client.Close()

	respUser, err := client.UpdateUser(user)

	if err != nil {
		t.Error(err.Error())
	}

	if respUser.GetName() != USER_MOCK_NAME {
		t.Error("not update user name")
	}

}
