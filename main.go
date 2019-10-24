package main

import (
	"log"
	"net/http"
	"os"
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

	// Any .env file with following parameters will be compatible;
	userName := os.Getenv("db_user")
	userPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	db, err := driver.ConnectPostgres(userName, userPass, dbName, dbHost, dbPort)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))

	centerHandler := handler.NewCenterHandler(db)
	r.Post("/centers", centerHandler.Create)
	r.Get("/centers", centerHandler.Fetch)
	//r.Get("/centers/{id}/courses", centerHandler.Fetch)
	//r.Get("/centers/{id}/departments", centerHandler.Fetch)

	depHandler := handler.NewDepartmentHandler(db)
	r.Post("/departments", depHandler.Create)
	r.Get("/departments", depHandler.Fetch)

	courseHandler := handler.NewCourseHandler(db)
	r.Post("/courses", courseHandler.Create)
	r.Get("/courses", courseHandler.Fetch)

	r.Route("/students", router.NewStudentRouter(db))

	log.Printf("NEGO API started on port %s.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func main() {
	port := "8081"
	Init(port)
}
