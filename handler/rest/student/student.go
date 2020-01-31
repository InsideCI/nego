package student

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/student"
)

// Student is a handler that wraps a Student repository.
type Student struct {
	repo repository.StudentRepository
}

// NewStudentHandler creates a new handler with a specific database engine.
func NewStudentHandler(db *driver.DB) *Student {
	return &Student{
		repo: student.NewStudentRepository(db),
	}
}

func (s *Student) Create(w http.ResponseWriter, r *http.Request) {
	var created int
	ctx := r.Context()
	students := ctx.Value("students").([]model.Student)

	created, err := s.repo.Create(students)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	w.Write([]byte(strconv.Itoa(created)))
}

func (s *Student) Fetch(w http.ResponseWriter, r *http.Request) {
	students, err := s.repo.Fetch()
	if err != nil {
		log.Println("[HANDLER]", err)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// FetchOne uses id param as primary key search
func (s *Student) FetchOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	registration, _ := strconv.Atoi(ctx.Value("id").(string))

	student, err := s.repo.FetchOne(registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(student)

}
