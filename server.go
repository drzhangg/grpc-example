package main

import (
	"context"
	pb "github.com/drzhangg/grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server,&SearchService{})

	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err:%v",err)
	}

	server.Serve(lis)
}
