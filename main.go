package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/InsideCI/nego/handler"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Init method first
func Init(port string) {
	handler := handler.NewHandler()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", handler.GetStudent).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	port := "8081"
	fmt.Printf("Listening on port: %s", port)
	Init(port)
}
