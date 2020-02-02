package rest

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/services"
	"log"
	"net/http"
	"strconv"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
)

// StudentController is a controllers that wraps a StudentRepository.
type StudentController struct {
	service *services.StudentService
}

// NewStudentController creates a new controllers with a specific database engine.
func NewStudentController(db *driver.DB) *StudentController {
	return &StudentController{
		service: services.NewStudentService(db),
	}
}

func (s *StudentController) Create(w http.ResponseWriter, r *http.Request) {
	var created int
	ctx := r.Context()
	students := ctx.Value("students").([]models.Student)

	created, err := s.service.Create(students)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	w.Write([]byte(strconv.Itoa(created)))
}

func (s *StudentController) Fetch(w http.ResponseWriter, r *http.Request) {
	students, err := s.service.Fetch()
	if err != nil {
		log.Println("[HANDLER]", err)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// FetchOne uses id param as primary key search
func (s *StudentController) FetchOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	registration, _ := strconv.Atoi(ctx.Value("id").(string))

	student_, err := s.service.FetchOne(registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(student_)

}
