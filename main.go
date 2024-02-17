package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"badinho.com/baas-connector/book"
	"badinho.com/baas-connector/connection"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

func main() {
	pool := connection.Pool()
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	book.RegisterBookServiceServer(grpcServer, &book.Server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// c, err := bookMethods.Create(bookMethods.WriteBook{
// 	Title:      "Perales Bio 2",
// 	Author:     "Perales",
// 	Price:      50,
// 	CategoryID: 1,
// })
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Book added: %v\n", c)

// books, err := bookMethods.FindAll()
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Books found: %v\n", books)

// book, err := bookMethods.FindByUID(books[1].UID)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Book found: %v\n", book)

// r, err := bookMethods.Update(book.ID, bookMethods.WriteBook{
// 	Title:      "Perales Master",
// 	Author:     book.Author,
// 	Price:      book.Price,
// 	CategoryID: book.Category.ID,
// })
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Book added: %v\n", r)
