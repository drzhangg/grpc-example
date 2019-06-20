package main

import (
	"context"
	"fmt"
	pz "github.com/drzhangg/grpc-example/proto/demo"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("conn grpc server failed:%v", err)
	}
	defer conn.Close()

	//初始化客户端
	c := pz.NewDemoClient(conn)

	//调用方法
	reqBody := new(pz.DemoRequest)
	reqBody.Name = "hehehe"
	r, err := c.RpcDemo(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Message)
}
