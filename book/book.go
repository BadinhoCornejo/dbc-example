package book

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Server struct {
	UnimplementedGBookServiceServer
}

func (s *Server) GetBooks(ctx context.Context, in *GBooksRequest) (*GBooksResponse, error) {
	books, err := FindAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Books found: %v\n", books)

	var response []*GBookResponse

	for _, v := range books {
		response = append(response, &GBookResponse{
			Book: &GBook{
				Id:        v.ID,
				Uid:       v.UID,
				Title:     v.Title,
				Author:    v.Author,
				Price:     v.Price,
				CreatedAt: v.CreatedAt.Format(time.DateTime),
				UpdatedAt: v.UpdatedAt.Format(time.DateTime),
			},
		})
	}

	return &GBooksResponse{
		Books: response,
	}, nil
}

func (s *Server) GetBookByUid(ctx context.Context, in *GBookByUidRequest) (*GBookResponse, error) {
	book, err := FindByUID(in.Uid)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Book found: %v\n", book)

	return &GBookResponse{
		Book: &GBook{
			Id:        book.ID,
			Uid:       book.UID,
			Title:     book.Title,
			Author:    book.Author,
			Price:     book.Price,
			CreatedAt: book.CreatedAt.Format(time.DateTime),
			UpdatedAt: book.UpdatedAt.Format(time.DateTime),
		},
	}, nil
}

func (s *Server) DeleteBook(ctx context.Context, in *GBookByIdRequest) (*GDeleteResponse, error) {
	affected, err := DeleteByID(in.Id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Affected: %v\n", affected)

	return &GDeleteResponse{
		Affected: affected,
	}, nil
}

func (s *Server) CreateBook(ctx context.Context, in *GCreateBookRequest) (*GBookResponse, error) {
	uid, err := Create(WriteBook{
		Title:      in.Body.Title,
		Author:     in.Body.Author,
		Price:      in.Body.Price,
		CategoryID: in.Body.CategoryId,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Created: %v\n", uid)

	book, err := FindByUID(uid)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Book created: %v\n", book)

	return &GBookResponse{
		Book: &GBook{
			Id:        book.ID,
			Uid:       book.UID,
			Title:     book.Title,
			Author:    book.Author,
			Price:     book.Price,
			CreatedAt: book.CreatedAt.Format(time.DateTime),
			UpdatedAt: book.UpdatedAt.Format(time.DateTime),
		},
	}, nil
}

func (s *Server) UpdateBook(ctx context.Context, in *GUpdateBookRequest) (*GBookResponse, error) {
	book, err := FindByUID(in.Uid)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Old Book: %v\n", book)

	now := time.Now().UTC().Format(time.DateTime)

	affected, err := Update(
		book.ID,
		WriteBook{
			Title:      in.Body.Title,
			Author:     in.Body.Author,
			Price:      in.Body.Price,
			CategoryID: in.Body.CategoryId,
		},
		now,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Affected: %v\n", affected)

	return &GBookResponse{
		Book: &GBook{
			Id:        book.ID,
			Uid:       book.UID,
			Title:     in.Body.Title,
			Author:    in.Body.Author,
			Price:     in.Body.Price,
			CreatedAt: book.CreatedAt.Format(time.DateTime),
			UpdatedAt: now,
		},
	}, nil
}
