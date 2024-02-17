package main

import (
	"context"
	"flag"
	"log"
	"time"

	"badinho.com/dbc-example/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := book.NewBookServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.GetBooks(ctx, &book.BooksRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetBooks: %s", err)
	}

	log.Printf("Response from server: %v", response)
}
