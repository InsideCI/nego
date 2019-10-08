package driver

import (
	"fmt"
	"log"

	"github.com/InsideCI/nego/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres dialect
)

// DB wrappes available projet databases.
type DB struct {
	Psql *gorm.DB
	//sqlserver *gorm.DB
}

// ConnectPostgres returns an instance of Postgres database connection.
func ConnectPostgres(user, pass, name, host, port string) (*DB, error) {

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", host, port, name, user, pass)
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL database; Starting migration..")
	db.AutoMigrate(&model.Center{}, &model.Department{})
	log.Println("Migration ended with no errors.")

	return &DB{
		Psql: db,
	}, nil
}
