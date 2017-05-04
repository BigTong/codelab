package rpcbm

import (
	"encoding/json"
	"flag"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/BigTong/common/log"
	"github.com/golang/protobuf/proto"
	"github.com/pquerna/ffjson/ffjson"
)

var grpcServerStarted = false
var mutex = &sync.Mutex{}

func init() {
	runtime.GOMAXPROCS(8)
	logConfig := flag.String("log_config", "conf/log.json", "")
	flag.Parse()

	log.InitFileLoggerFromConfigFile(*logConfig, log.INFO)
}

func StartServer() {
	mutex.Lock()
	defer mutex.Unlock()

	if grpcServerStarted {
		return
	}
	go MockGrpcServer()
	go MockJsonRpcServer()
	time.Sleep(2 * time.Second)
	grpcServerStarted = true
}

func BenchmarkGrpcUpdateUser(b *testing.B) {
	StartServer()

	b.StopTimer()

	client := NewGrpcClient(GRPC_SERVER_ADDRESS)
	defer client.Close()

	user := NewFakeProtoUser()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		client.UpdateUser(user)
	}
}

func BenchmarkJsonrpcUpdateUser(b *testing.B) {
	StartServer()

	b.StopTimer()

	client := NewJsonRpcClient(JSONRPC_SERVER_ADDRESS, 5)
	user := NewFakeUser()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		client.UpdateUser(user)
	}
}

func BenchmarkJsonSerialAndDeSelrial(b *testing.B) {
	b.StopTimer()
	user := NewFakeUser()
	userData, _ := json.Marshal(user)
	log.FInfo("get json len:%d", len(userData))

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(user)

		json.Unmarshal(data, user)
	}

}

func BenchmarkFJsonSerialAndDeSelrial(b *testing.B) {
	b.StopTimer()
	user := NewFakeUser()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data, _ := ffjson.Marshal(user)
		ffjson.Unmarshal(data, user)
	}

}

func BenchmarkProtoSerialAndDeSelrial(b *testing.B) {
	b.StopTimer()
	user := NewFakeProtoUser()
	userData, _ := proto.Marshal(user)
	log.FInfo("get proto len:%d", len(userData))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data, _ := proto.Marshal(user)
		proto.Unmarshal(data, user)
	}

}
