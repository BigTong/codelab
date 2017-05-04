package rpcbm

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/BigTong/common/log"
)

const (
	JSONRPC_SERVER_ADDRESS = "localhost:9091"
)

type JsonRpcServer struct {
}

func MockJsonRpcServer() {
	server := rpc.NewServer()
	listener, err := net.Listen("tcp", JSONRPC_SERVER_ADDRESS)
	if err != nil {
		log.FFatal("listen get error:%s", err.Error())
	}
	defer listener.Close()

	handler := &JsonRpcServer{}
	server.Register(handler)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.FFatal("listen get err:%s", err.Error())
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func (self *JsonRpcServer) UpdateUser(user *User, resp *User) error {
	// ToDo(add deepcopy)
	*resp = *user
	resp.Name = USER_MOCK_NAME
	log.FInfo("get user:%s and set user name:%s",
		user.Name, USER_MOCK_NAME)
	return nil
}
