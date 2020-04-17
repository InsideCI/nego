package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/router"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {

	port := flag.String("port", constants.DefaultPort, "API port.")
	debug := flag.Bool("debug", constants.DefaultFlagDebug, "SQL debug switch.")
	flag.Parse()

	if err := godotenv.Load("app.env"); err != nil {
		panic("You must provide a .env config file. Instructions at README.")
	}
	//TODO: check is JWT key was set.

	// DATABASE
	db, err := driver.CreateDatabasesConnections(*debug)
	if err != nil {
		panic("Could not create database connection: " + err.Error())
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
