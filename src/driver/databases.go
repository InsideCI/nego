package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/InsideCI/nego/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres dialect
)

// DB wrappes available projet databases.
type DB struct {
	Postgres *gorm.DB
	//Mongodb *gorm.DB
}

// CreateDatabasesConnections returns an instance of predefined
// databases connections.
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
		return nil, err
	}

	log.Println("Connected to Postgres database. Starting migration...")

	validateAndMigrate(db,
		&models.Center{},
		&models.Department{},
		&models.Teacher{},
		&models.Class{},
		&models.Course{},
		&models.Student{},
		&models.GeneralStatistic{})

	log.Println("Migration ended with no errors.")

	return &DB{
		Postgres: db,
	}, nil
}

//validateAndMigrate checks if all models implement Nego interface in compile time.
func validateAndMigrate(db *gorm.DB, models ...models.Nego) {
	for _, v := range models {
		db.AutoMigrate(v)
	}
}
