package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/mobiledatabooks/gcp-go-grpc/mobiledatabooks.com/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func callSayHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		fmt.Println("SayHello: _, ", err)
	} else {
		fmt.Println("SayHello: ", r.GetMessage())
	}
}
func main() {
	flag.Parse()
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // block until connection is established
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, options...)
	if err != nil { // if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	greeterClient := pb.NewGreeterClient(conn) // Create a client

	counter := 0
	for { // loop forever
		start := time.Now()
		callSayHello(greeterClient)                                      // call the function
		time.Sleep(1 * time.Second)                                      // sleep for 1 second
		fmt.Println("It took ", time.Since(start), " counter:", counter) // print the time it took
		counter++
	}

}
