package microservice

import (
	"context"
	"flag"
	"log"
	"time"

	pb "mytelegrambot/microservice/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName  = "list"
	defaultName2 = "world"
)

var (
	addr             = flag.String("addr", "localhost:50051", "the address to connect to")
	user_input       = flag.String("user_input", defaultName, "Name to greet")
	task_description = flag.String("task_description", defaultName, "Name to greet")
)

func Microconnection() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{UserInput: *user_input, TaskDescription: *task_description})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
