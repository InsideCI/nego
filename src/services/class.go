package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type ClassService struct {
	repo *repositories.ClassRepository
}

func NewClassService() *ClassService {
	return &ClassService{
		repo: repositories.NewClassRepository(),
	}
}

func (s *ClassService) Create(db *driver.DB, class models.Class) (*models.Class, error) {
	err := s.repo.Create(db.Postgres, &class)
	return &class, err
}

func (s *ClassService) Fetch(db *driver.DB, params models.QueryParams, example models.Class) (*models.Page, error) {
	return s.repo.FetchWithPagination(db.Postgres, params, example)
}

func (s *ClassService) FetchOne(db *driver.DB, id string) (*models.Class, error) {
	temp, err := s.repo.FetchOne(db.Postgres, id)
	return temp.(*models.Class), err
}
