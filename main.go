package main

import (
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/handler/rest/center"
	"github.com/InsideCI/nego/handler/rest/course"
	"github.com/InsideCI/nego/handler/rest/department"
	"github.com/InsideCI/nego/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// Init method first
func Init(port string) {

	// Database init configurations
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	dbConnection, err := driver.CreateDatabasesConnections()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	//r.Use(middleware.Timeout(5 * time.Second))

	centerHandler := center.NewCenterHandler(dbConnection)
	r.Post("/centers", centerHandler.Create)
	r.Get("/centers", centerHandler.Fetch)

	depHandler := department.NewDepartmentHandler(dbConnection)
	r.Post("/departments", depHandler.Create)
	r.Get("/departments", depHandler.Fetch)

	courseHandler := course.NewCourseHandler(dbConnection)
	r.Post("/courses", courseHandler.Create)
	r.Get("/courses", courseHandler.Fetch)

	r.Route("/students", router.NewStudentRouter(dbConnection))

	log.Printf("NEGO API started on port %s.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func main() {
	port := "8081"
	Init(port)
}
