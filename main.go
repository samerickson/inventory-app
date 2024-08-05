package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
}

	// Define the connection string.
	user:= os.Getenv("POSTGRES_USER")
	pass:= os.Getenv("POSTGRES_PASSWORD")

	connStr := "user=" + user + " dbname=default_database sslmode=disable password=" + pass

	// Open the connection.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify the connection.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}
