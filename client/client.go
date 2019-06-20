package main

import (
	"context"
	pb "github.com/drzhangg/grpc-example/proto"
	"github.com/gpmgo/gopm/modules/log"
	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc.Dial err:%v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatal("client.Search err:%v", err)
	}
	log.Fatal("resp:%s", resp.GetResponse())

}
