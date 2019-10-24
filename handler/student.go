package handler

import (
	"encoding/json"
	"fmt"
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
	var student model.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		log.Println(err)
		return
	}
	s.repo.Create(&student)
	w.Write([]byte("OK"))
}

func (s *Student) Fetch(w http.ResponseWriter, r *http.Request) {
	students, err := s.repo.Fetch()
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// FetchOne uses id param as primary key search
func (s *Student) FetchOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("id")
	reg, err := strconv.ParseInt(fmt.Sprintf("%v", id), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	student, err := s.repo.FetchOne(reg)
	err = json.NewEncoder(w).Encode(student)
}
