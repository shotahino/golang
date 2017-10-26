package main

import (
	"log"
	"os"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "192.168.7.105:3000"
	defaultName = "Hino"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMOTDClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Greeting(context.Background(), &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.Motd)
}
