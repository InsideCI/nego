package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type TeacherService struct {
	repo *repositories.TeacherRepository
}

func NewTeacherService() *TeacherService {
	return &TeacherService{
		repo: repositories.NewTeacherRepository(),
	}
}

func (s *TeacherService) Create(db *driver.DB, teacher models.Teacher) (*models.Teacher, error) {
	err := s.repo.Create(db.Postgres, &teacher)
	return &teacher, err
}

func (s *TeacherService) Fetch(db *driver.DB, teacher models.Teacher) (*models.Teacher, error) {
	err := s.repo.Create(db.Postgres, &teacher)
	return &teacher, err
}

func (s *TeacherService) FetchOne(db *driver.DB, id string) (*models.Student, error) {
	temp, err := s.repo.FetchOne(db.Postgres, id)
	return temp.(*models.Student), err
}
