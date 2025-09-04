package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//dbDriver := os.Getenv("dbDriver")
	dbSource := os.Getenv("dbSource")

	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("couldn't connect to the database.")
	}
	//defer pool.Close()
	testQueries = New(pool)
	code := m.Run()
	os.Exit(code)

}
