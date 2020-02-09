package rest

import (
	"encoding/json"
	"fmt"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/services"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"log"
	"net/http"
)

// StudentController is a controller that wraps a StudentRepository.
type StudentController struct {
	db      *driver.DB
	service *services.StudentService
}

// NewStudentController creates a new controller with a specific database engine.
func NewStudentController(db *driver.DB) *StudentController {
	return &StudentController{
		db:      db,
		service: services.NewStudentService(),
	}
}

func (s *StudentController) Create(w http.ResponseWriter, r *http.Request) {
	student := r.Context().Value("payload").(model.Student)

	created, err := s.service.Create(s.db, student)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (s *StudentController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(map[string][]string)

	page, err := s.service.Fetch(s.db, params)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)
}

// FetchOne uses id param as primary key search
func (s *StudentController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	student, err := s.service.FetchOne(s.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
	}

	_ = json.NewEncoder(w).Encode(student)

}
