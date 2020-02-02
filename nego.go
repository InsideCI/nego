package main

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/routers"
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

	//centerHandler := center.NewCenterHandler(dbConnection)
	//r.Post("/centers", centerHandler.Create)
	//r.Get("/centers", centerHandler.Fetch)

	r.Route("/students", routers.NewStudentRouter(dbConnection))

	log.Printf("NEGO API started on port %s.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func main() {
	port := "8081"
	Init(port)
}
