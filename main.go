package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// todo: better error handling for env variables
	password, passwordErr := os.LookupEnv("DB_PASSWORD")
	hostname, _ := os.LookupEnv("DB_HOSTNAME")
	database, _ := os.LookupEnv("DB_DATABASE")
	user, _ := os.LookupEnv("DB_USER")

	if passwordErr != true {
		log.Fatal("Error looking up database password", passwordErr)
	}
	connStr := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", user, password, hostname, database)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	var res string
	queryErr := db.QueryRow("SELECT title FROM books WHERE id = 1").Scan(&res)
	if queryErr != nil {
		fmt.Println(queryErr)
	}

	fmt.Println(res)
}
