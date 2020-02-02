package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(db *driver.DB) *StudentService {
	return &StudentService{
		repo: repositories.NewStudentRepository(),
	}
}

func (s *StudentService) Create(student models.Student) interface{} {
	return s.repo.Create(&student)
}
