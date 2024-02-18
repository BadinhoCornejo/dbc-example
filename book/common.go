package book

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"badinho.com/dbc-example/connection"
	"github.com/google/uuid"
)

func FindAll() ([]Book, error) {
	conn := connection.Pool()
	var books []Book

	rows, err := conn.Query(context.Background(), "select b.id, b.uid, b.title, b.author, b.price, b.created_at, b.updated_at, c.id, c.uid, c.name, c.created_at, c.updated_at from book b inner join category c on(b.category_id = c.id) where b.status = B'1'")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(
			&book.ID,
			&book.UID,
			&book.Title,
			&book.Author,
			&book.Price,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Category.ID,
			&book.Category.UID,
			&book.Category.Name,
			&book.Category.CreatedAt,
			&book.Category.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("books %v", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("books %v", err)
	}

	return books, nil
}

func FindByUID(uid string) (Book, error) {
	conn := connection.Pool()
	var book Book

	err := conn.QueryRow(
		context.Background(),
		fmt.Sprintf("select b.id, b.uid, b.title, b.author, b.price, b.created_at, b.updated_at, c.id, c.uid, c.name, c.created_at, c.updated_at from book b inner join category c on(b.category_id = c.id) where b.uid = '%s' and b.status = B'1'", uid),
	).Scan(
		&book.ID,
		&book.UID,
		&book.Title,
		&book.Author,
		&book.Price,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.Category.ID,
		&book.Category.UID,
		&book.Category.Name,
		&book.Category.CreatedAt,
		&book.Category.UpdatedAt,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return book, fmt.Errorf("book %v", err)
	}

	return book, nil
}

func Create(book WriteBook) (string, error) {
	uid := uuid.NewString()
	conn := connection.Pool()
	now := time.Now().UTC().Format(time.DateTime)

	_, err := conn.Exec(
		context.Background(),
		"insert into book (uid, title, author, price, created_at, updated_at, category_id, status) values ($1, $2, $3, $4, $5, $6, $7, B'1')",
		uid,
		book.Title,
		book.Author,
		book.Price,
		now,
		now,
		book.CategoryID,
	)
	if err != nil {
		return "", fmt.Errorf("create book: %w", err)
	}

	return uid, nil
}

func Update(id int64, book WriteBook, now string) (int64, error) {
	conn := connection.Pool()

	result, err := conn.Exec(
		context.Background(),
		"update book set title = $2, author = $3, price = $4, updated_at = $5, category_id = $6 where id = $1 and status = B'1'",
		id,
		book.Title,
		book.Author,
		book.Price,
		now,
		book.CategoryID,
	)
	if err != nil {
		return 0, fmt.Errorf("update book: %w", err)
	}

	log.Default().Println(result.String())

	return result.RowsAffected(), nil
}

func DeleteByID(id int64) (int64, error) {
	conn := connection.Pool()
	now := time.Now().UTC().Format(time.DateTime)

	result, err := conn.Exec(
		context.Background(),
		"update book set status = B'0', updated_at = $2 where id = $1",
		id,
		now,
	)
	if err != nil {
		return 0, fmt.Errorf("update book: %w", err)
	}

	log.Default().Println(result.String())

	return result.RowsAffected(), nil
}
