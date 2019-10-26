package main

import (
	"log"
	"net/http"
	"time"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/handler"
	"github.com/InsideCI/nego/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

// Init method first
func Init(port string) {

	// Database init configurations
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	databasesConnection, err := driver.CreateDatabasesConnections()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))

	centerHandler := handler.NewCenterHandler(databasesConnection)
	r.Post("/centers", centerHandler.Create)
	r.Get("/centers", centerHandler.Fetch)

	depHandler := handler.NewDepartmentHandler(databasesConnection)
	r.Post("/departments", depHandler.Create)
	r.Get("/departments", depHandler.Fetch)

	courseHandler := handler.NewCourseHandler(databasesConnection)
	r.Post("/courses", courseHandler.Create)
	r.Get("/courses", courseHandler.Fetch)

	r.Route("/students", router.NewStudentRouter(databasesConnection))

	log.Printf("NEGO API started on port %s.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func main() {
	port := "8081"
	Init(port)
}
