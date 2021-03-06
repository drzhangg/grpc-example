package main

import (
	"context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const(
	port = ":5051"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context,in *pb.HelloRequest) (*pb.HelloReply,error) {
	return &pb.HelloReply{Message:"hello " + in.Name},nil
}
func main() {
	lis,err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
