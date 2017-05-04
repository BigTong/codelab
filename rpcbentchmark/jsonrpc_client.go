package rpcbm

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"github.com/BigTong/common/log"
)

type JsonRpcClient struct {
	conn   net.Conn
	client *rpc.Client
}

func NewJsonRpcClient(address string, timeout int) *JsonRpcClient {
	ret := &JsonRpcClient{}
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		log.FFatal("get err:%s", err.Error())
	}
	ret.conn = conn
	ret.client = jsonrpc.NewClient(conn)
	return ret
}

func (self *JsonRpcClient) Close() {
	if self.client != nil {
		self.client.Close()
	}
}

func (self *JsonRpcClient) UpdateUser(user *User) (*User, error) {
	respUser := &User{}
	err := self.client.Call("JsonRpcServer.UpdateUser", user, respUser)
	if err != nil {
		log.FError("call get error:%s", err.Error())
		return nil, err
	}
	return respUser, nil
}
