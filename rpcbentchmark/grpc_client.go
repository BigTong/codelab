package rpcbm

import (
	"time"

	"codelab/rpcbentchmark/proto"
	"github.com/BigTong/common/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	MAX_RETRY_COUNT = 3
)

func GetDialOptions() []grpc.DialOption {
	ret := []grpc.DialOption{}
	ret = append(ret, grpc.WithTimeout(20*time.Second))
	ret = append(ret, grpc.WithBlock())
	ret = append(ret, grpc.WithInsecure())
	return ret
}

func NewGrpcClient(address string) *GrpcClient {
	ret := &GrpcClient{
		ServerAddress: address,
	}

	option := GetDialOptions()
	conn, err := grpc.Dial(address, option...)
	if err != nil {
		log.FError("dial server %s get error:%s", address, err.Error())
		return nil
	}
	ret.Conn = conn

	log.FInfo("dial server ok:%s", address)

	ret.client = proto.NewUserServerClient(conn)
	return ret
}

type GrpcClient struct {
	ServerAddress string
	Conn          *grpc.ClientConn

	client proto.UserServerClient
}

func (gc *GrpcClient) Close() {
	if gc.Conn != nil {
		gc.Conn.Close()
	}
}

func (gc *GrpcClient) UpdateUser(user *proto.User) (*proto.User, error) {
	var respUser *proto.User
	var err error
	for i := 0; i < MAX_RETRY_COUNT; i++ {
		respUser, err = gc.client.UserUpdate(context.Background(), user)
		if err == nil {
			return respUser, nil
		}
		time.Sleep(100 * time.Millisecond)
		gc.Redial()
	}
	return respUser, err

}

func (gc *GrpcClient) Redial() bool {
	option := GetDialOptions()
	conn, err := grpc.Dial(gc.ServerAddress, option...)
	if err != nil {
		log.EError("dial server %s get error:%s", gc.ServerAddress, err.Error())
		return false
	}
	gc.Conn = conn
	gc.client = proto.NewUserServerClient(conn)
	return true
}
