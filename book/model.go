package book

import "time"

type Book struct {
	ID        int64
	UID       string
	Title     string
	Author    string
	Price     float32
	CreatedAt time.Time
	UpdatedAt time.Time
	Category  Category
}

type Category struct {
	ID        int64
	UID       string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WriteBook struct {
	Title      string
	Author     string
	Price      float32
	CategoryID int64
}
