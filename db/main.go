package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func goDotEnvVariable(key string) string {

	if godotenv.Load(".env") != nil {
		if godotenv.Load("/etc/secrets/.env") != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return os.Getenv(key)
}

func ConnectToDatabase() *sql.DB {
	var err error
	pUser := goDotEnvVariable("POSTGRES_USER")
	pPass := goDotEnvVariable("POSTGRES_PASSWORD")
	pDB := goDotEnvVariable("POSTGRES_DB")
	pHost := goDotEnvVariable("POSTGRES_HOST")

	urlDB := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", pHost, pUser, pPass, pDB)

	db, err = sql.Open("postgres", urlDB)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to yout database.")

	return db
}
