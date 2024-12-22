package dbSetup

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDBConnection() *pgxpool.Pool {
	// getting all variables from env
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// preparing connection string
	connectionString := "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + ""

	// creating a connection pool
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		fmt.Printf("Unable to parse connection string: %v\n", err)
	}
	// connecting to database
	session, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		panic(err)
	}

	// testing the connection with greeting query
	var greeting string
	err = session.QueryRow(context.Background(), "SELECT 'Hello, PostgreSQL!'").Scan(&greeting)
	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
		panic(err)
	}
	fmt.Println("Connected to PostgreSQL!")

	return session
}
