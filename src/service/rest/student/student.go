package student

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/repository"
	"github.com/InsideCI/nego/src/repository/student"
)

// StudentService is a service that wraps a StudentRepository.
type StudentService struct {
	repo repository.StudentRepository
}

// NewStudentService creates a new service with a specific database engine.
func NewStudentService(db *driver.DB) *StudentService {
	return &StudentService{
		repo: student.NewStudentRepository(db),
	}
}

func (s *StudentService) Create(w http.ResponseWriter, r *http.Request) {
	var created int
	ctx := r.Context()
	students := ctx.Value("students").([]model.Student)

	created, err := s.repo.Create(students)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	w.Write([]byte(strconv.Itoa(created)))
}

func (s *StudentService) Fetch(w http.ResponseWriter, r *http.Request) {
	students, err := s.repo.Fetch()
	if err != nil {
		log.Println("[HANDLER]", err)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// FetchOne uses id param as primary key search
func (s *StudentService) FetchOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	registration, _ := strconv.Atoi(ctx.Value("id").(string))

	student_, err := s.repo.FetchOne(registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(student_)

}
