package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/repository"
)

type TeacherService struct {
	repo *repository.TeacherRepository
}

func NewTeacherService() *TeacherService {
	return &TeacherService{
		repo: repository.NewTeacherRepository(),
	}
}

func (s *TeacherService) Create(db *driver.DB, teacher model.Teacher) (*model.Teacher, error) {
	err := s.repo.Create(db.Postgres, &teacher)
	return &teacher, err
}

func (s *TeacherService) Fetch(db *driver.DB, teacher model.Teacher) (*model.Teacher, error) {
	err := s.repo.Create(db.Postgres, &teacher)
	return &teacher, err
}

func (s *TeacherService) FetchOne(db *driver.DB, id string) (*model.Student, error) {
	temp, err := s.repo.FetchOne(db.Postgres, id)
	return temp.(*model.Student), err
}
