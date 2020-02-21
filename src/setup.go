package setup

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/router"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
)

func Init() {

	// DATABASE
	dbConnection, err := driver.CreateDatabasesConnections()
	if err != nil {
		panic(err)
	}

	// CONTROLLER
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middlewares.Cors.Handler)
	router.InitRoutes(dbConnection, r)

	port := os.Getenv("api_port")
	log.Printf("NEGO API started on port %s.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
