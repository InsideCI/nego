package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/InsideCI/nego/store"
)

// Handler defines a simple Router struct
type Handler struct {
	st *store.Store
}

// NewHandler creates an instante of a Handler containing a database connection.
func NewHandler() *Handler {
	return &Handler{
		st: store.NewStore(),
	}
}

// GetStudent returns an instante of a student by it's registration code.
func (ro *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>GET STUDENT ENDPOINT WORKING.</h1>")
}

// GetCenters returns all centers as JSON objects
func (ro *Handler) GetCenters(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>GET CENTERS ENDPOINT WORKING.</h1>")
	fmt.Println("Calling store method.")
	centers := ro.st.GetCenters()
	json.NewEncoder(w).Encode(centers)
}
