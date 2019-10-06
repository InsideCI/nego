package store

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Store abstracts CRUD methods
type Store struct {
	db *gorm.DB
}

// NewStore creates and returns a database based on .env file.
func NewStore() *Store {

	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	userName := os.Getenv("db_user")
	userPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s dbname=%s user=%s password=%s  sslmode=disable", dbHost, dbName, userName, userPassword)

	db, err := gorm.Open("postgres", dbURI)

	return &Store{
		db: db,
	}
}

// GetStudent returns a student name and course ID based on it's registration code.
// func (s *Store) GetStudent() (*model.Student, error) {
// 	return _, _
// }
