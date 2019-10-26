package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/InsideCI/nego/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres dialect
)

// DB wrappes available projet databases.
type DB struct {
	Psql *gorm.DB
	//sqlserver *gorm.DB
}

// CreateDatabasesConnections returns an instance of Postgres database connection.
func CreateDatabasesConnections() (*DB, error) {

	// Any .env file with following parameters will be compatible;
	user := os.Getenv("db_user")
	pass := os.Getenv("db_pass")
	name := os.Getenv("db_name")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", host, port, name, user, pass)
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres database. Starting migration...")
	db.AutoMigrate(&model.Center{}, &model.Department{}, &model.Course{}, &model.Student{})
	log.Println("Migration ended with no errors.")

	return &DB{
		Psql: db,
	}, nil
}
