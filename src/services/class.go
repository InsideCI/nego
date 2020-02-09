package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/repository"
)

type ClassService struct {
	repo *repository.ClassRepository
}

func NewClassService() *ClassService {
	return &ClassService{
		repo: repository.NewClassRepository(),
	}
}

func (s *ClassService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}
