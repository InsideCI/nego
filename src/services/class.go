package services

import (
	"github.com/InsideCI/nego/src/driver"
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

func (s *ClassService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}
