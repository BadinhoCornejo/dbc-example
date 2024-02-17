package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool = nil

func Pool() *pgxpool.Pool {
	if dbpool != nil {
		return dbpool
	}

	newPool, err := pgxpool.New(context.Background(), "postgresql://postgres:password@localhost:5432/books")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	dbpool = newPool

	log.Default().Println("New Connection")

	return dbpool
}
