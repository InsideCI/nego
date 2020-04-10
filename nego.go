package main

import (
	"flag"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/router"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	port := flag.String("port", "8080", "API port.")
	debug := flag.Bool("debug", false, "SQL debug switch.")
	flag.Parse()

	if err := godotenv.Load("app.env"); err != nil {
		panic("You must provide a .env config file. Instructions at README.")
	}

	// DATABASE
	db, err := driver.CreateDatabasesConnections(*debug)
	if err != nil {
		panic("Could'nt create database connection: " + err.Error())
	}

	// CONTROLLER
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middlewares.Cors.Handler)
	router.InitRoutes(db, r)

	certificate := os.Getenv("certificate")
	key := os.Getenv("key")

	log.Printf("Starting NEGO at port %s.\n", *port)
	if certificate != "" && key != "" {
		log.Fatal(http.ListenAndServeTLS(":"+*port, certificate, key, r))
	} else {
		log.Fatal(http.ListenAndServe(":"+*port, r))
	}
}
