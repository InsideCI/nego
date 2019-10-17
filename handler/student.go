package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/student"
)

// Student is a handler that wraps a Student repository.
type Student struct {
	repo repository.StudentRepository
}

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
	students, err := s.repo.Fetch(10)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(students)
}
