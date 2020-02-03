package rest

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/services"
	"github.com/go-chi/chi"
	"log"
	"net/http"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
)

// StudentController is a controllers that wraps a StudentRepository.
type StudentController struct {
	db      *driver.DB
	service *services.StudentService
}

// NewStudentController creates a new controllers with a specific database engine.
func NewStudentController(db *driver.DB) *StudentController {
	return &StudentController{
		db:      db,
		service: services.NewStudentService(),
	}
}

func (s *StudentController) Fetch(w http.ResponseWriter, r *http.Request) {
	ctx := chi.Context{}
	students, err := s.service.Fetch()
	if err != nil {
		log.Println("[HANDLER]", err)
		return
	}
	json.NewEncoder(w).Encode(students)
}
