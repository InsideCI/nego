package rest

import (
	"encoding/json"
	"fmt"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/services"
	"log"
	"net/http"
	"strconv"
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

func (s *StudentController) Create(w http.ResponseWriter, r *http.Request) {
	student := r.Context().Value("student").(models.Student)

	created, err := s.service.Create(s.db, student)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	json.NewEncoder(w).Encode(created)
}

func (s *StudentController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(map[string][]string)

	students, err := s.service.Fetch(s.db, params)
	if err != nil {
		//log.Println("[HANDLER]", err)
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// FetchOne uses id param as primary key search
func (s *StudentController) FetchOne(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value("id").(string)
	registration, _ := strconv.Atoi(value)

	student, err := s.service.FetchOne(s.db, registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(student)

}
