package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Init method first
func Init(port string) {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	userName := os.Getenv("db_user")
	userPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	db, err := driver.ConnectPostgres(userName, userPass, dbName, dbHost, dbPort)
	if err != nil {
		log.Fatal(err)
	}

	centerHandler := handler.NewCenterHandler(db)
	r.Post("/centers", centerHandler.Create)
	r.Get("/centers", centerHandler.Fetch)

	//depHandler := handler.NewDepHandler()

	http.ListenAndServe(":"+port, r)
}

func main() {
	port := "8081"
	log.Printf("Starting server on port: %s\n", port)
	Init(port)
}
