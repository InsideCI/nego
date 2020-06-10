package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/router"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	port := flag.String("port", constants.DefaultPort, "API port.")
	debug := flag.Bool("debug", constants.DefaultFlagDebug, "SQL debug switch.")
	prod := flag.Bool("prod", constants.DefaultFlagTLS, "SQL debug switch.")
	flag.Parse()

	if err := godotenv.Load("app.env"); err != nil {
		panic("You must provide a .env config file. Instructions at README.")
	}
	//TODO: check if JWT key was set.

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

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("cert-cache"),
		HostPolicy: autocert.HostWhitelist("cinside.ddns.net"),
	}

	server := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	if *prod {
		log.Printf("Starting NEGO with a TLS connection at port %s.\n", *port)
		go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		log.Fatal(server.ListenAndServeTLS("", ""))
	}

	log.Printf("Starting NEGO with non TLS connection at port %s.\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))

}
