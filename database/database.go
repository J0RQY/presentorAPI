package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool
var once sync.Once

func InitDB() {
	once.Do(func() {
		connStr := os.Getenv("DATABASE_URL")
		pool, err := pgxpool.New(context.Background(), connStr)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v\n", err)
		}

		var greeting string
		err = pool.QueryRow(context.Background(), "select 'Database connection successful'").Scan(&greeting)
		if err != nil {
			log.Fatalf("Database connection check failed: %v\n", err)
		}
		db = pool
	})
}

func GetDB() *pgxpool.Pool {
	if db == nil {
		log.Fatal("Database connection not initialized. Ensure InitDB() is called once at application startup.")
	}
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection pool closed.")
	}
}
