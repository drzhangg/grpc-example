package main

import (
	"context"
	"fmt"
	pz "github.com/drzhangg/grpc-example/proto/demo"
	"google.golang.org/grpc"
	"net"
)

const (
	Address = "127.0.0.1:50052"
)

type DemoServer struct {
}

func (d *DemoServer) RpcDemo(ctx context.Context, in *pz.DemoRequest) (*pz.DemoReply, error) {
	resp := new(pz.DemoReply)
	resp.Message = "hello, this is " + in.Name + "."
	return resp, nil
}

var demoServer = DemoServer{}

func main() {

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}

	//实现grpc server
	s := grpc.NewServer()
	//注册demoserver为客户端提供服务
	pz.RegisterDemoServer(s, &demoServer)	//内部调用s.Register()
	fmt.Println("Listen on :", Address)

	s.Serve(listen)
}
