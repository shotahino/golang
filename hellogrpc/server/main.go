package main

import (
	"log"
	"net"
	"os"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement hellogrpc.GreeterServer.
type server struct{}

// SayHello implements hellogrpc.MOTDServer
func (s *server) Greeting(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Motd: "Dummy Service --- Welcome " + in.Name}, nil
}

func main() {
	log.SetOutput(os.Stdout)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMOTDServer(s, &server{})
	log.Println("gRPC server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
