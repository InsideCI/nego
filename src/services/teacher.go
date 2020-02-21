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
		repo: &repository.TeacherRepository{struct{ Type interface{} }{Type: model.Teacher{}}},
	}
}

func (s *TeacherService) Create(db *driver.DB, teacher model.Teacher) (*model.Teacher, error) {
	temp, err := s.repo.Create(db.Postgres, &teacher)
	teacher = temp.(model.Teacher)
	return &teacher, err
}
