package main

import (
	"context"
	"log"
	pd "source/grpc/proto"
	"time"

	"google.golang.org/grpc"
)

var (
	address     = "localhost:8080"
	defaultName = "Your name"
)

func main() {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	c := pd.NewGreetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.Hello(ctx, &pd.SayHello{Name: defaultName})
	r2, _ := c.HelloAgain(ctx, &pd.SayHello{Name: defaultName})
	log.Println("Greeting: ", r.GetMessage())
	log.Println("Greeting: ", r2.GetMessage())
}
