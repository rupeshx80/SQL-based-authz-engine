package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Pool *pgxpool.Pool

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("dbURL is not set")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("failed to ping db:", err)
	}

	Pool = pool
	log.Println("postgres connected")
}
