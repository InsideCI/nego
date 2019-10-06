package routes

import (
	"fmt"
	"net/http"

	"github.com/InsideCI/nego/store"
)

// Router defines a simple Router struct
type Router struct {
	db *store.Database
}

// NewRouter creates an instante of a router containing a database connection.
func NewRouter() *Router {
	return &Router{
		db: store.NewDatabase(),
	}
}

// GetStudent returns an instante of a student by it's registration code.
func GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users endpoint hit.")
}
