package main

import (
	"context"
	"log"
	"time"

	pb "actor-grpc-demo/proto/proto-demo/myapp/helloworldpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Adam"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Greeting:", res.Message)
}
