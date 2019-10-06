package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Init method first
func Init(port string) {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//centerHandler := handler.NewCenterHandler()
	// r.Get("/centers", handler.Fetch)

	//depHandler := handler.NewDepHandler()

	http.ListenAndServe(":"+port, r)
}

func main() {
	port := "8081"
	fmt.Printf("Listening on port: %s\n", port)
	Init(port)
}
