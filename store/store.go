package store

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	userName := os.Getenv("db_user")
	userPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s dbname=%s user=%s password=%s  sslmode=disable", dbHost, dbName, userName, userPassword)

	db, err := gorm.Open("postgres", dbUri)

	return &Database{
		db: db,
	}
}
