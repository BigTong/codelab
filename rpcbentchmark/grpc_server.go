package rpcbm

import (
	"net"

	"codelab/rpcbentchmark/proto"

	"github.com/BigTong/common/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	USER_MOCK_NAME      = "lisi"
	GRPC_SERVER_ADDRESS = "localhost:9090"
)

func MockGrpcServer() {
	rpcLis, err := net.Listen("tcp", GRPC_SERVER_ADDRESS)
	if err != nil {
		log.FFatal("rpc get error:%s", err.Error())
	}

	userServer := NewGrpcServer()
	grpcServer := grpc.NewServer()
	proto.RegisterUserServerServer(grpcServer, userServer)

	go grpcServer.Serve(rpcLis)
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

type GrpcServer struct{}

func (self *GrpcServer) UserUpdate(ctx context.Context,
	user *proto.User) (*proto.User, error) {
	log.FInfo("get user:%s and set user name:%s",
		user.Name, USER_MOCK_NAME)
	user.Name = USER_MOCK_NAME
	return user, nil
}
