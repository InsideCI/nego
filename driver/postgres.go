package driver

import (
	"fmt"
	"log"

	"github.com/InsideCI/nego/model"
	"github.com/jinzhu/gorm"
)

// DB wrappes available projet databases.
type DB struct {
	psql *gorm.DB
	//sqlserver *gorm.DB
}

// ConnectPostgres returns an instance of Postgres database connection.
func ConnectPostgres(user, pass, name, host, port string) (*DB, error) {

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", host, port, name, user, pass)
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Center{})

	return &DB{
		psql: db,
	}, nil
}
