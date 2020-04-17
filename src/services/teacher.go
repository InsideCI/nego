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

func (s *TeacherService) Fetch(db *driver.DB, params models.QueryParams, example models.Teacher) (*models.Page, error) {
	return s.repo.FetchWithPagination(db.Postgres, params, example)
}

func (s *TeacherService) FetchOne(db *driver.DB, id string) (*models.Teacher, error) {
	temp, err := s.repo.FetchOne(db.Postgres, id)
	if err != nil {
		return nil, err
	}
	return temp.(*models.Teacher), err
}
