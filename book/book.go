package book

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Server struct {
	UnimplementedBookServiceServer
}

func (s *Server) GetBooks(ctx context.Context, in *BooksRequest) (*BooksResponse, error) {
	books, err := FindAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Books found: %v\n", books)

	var response []*BookResponse

	for _, v := range books {
		response = append(response, &BookResponse{
			Id:        v.ID,
			Uid:       v.UID,
			Title:     v.Title,
			Author:    v.Author,
			Price:     v.Price,
			CreatedAt: v.CreatedAt.Format(time.DateTime),
			UpdatedAt: v.UpdatedAt.Format(time.DateTime),
		})
	}

	return &BooksResponse{
		Books: response,
	}, nil
}
