package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type CenterService struct {
	repo *repositories.CenterRepository
}

func NewCenterService() *CenterService {
	return &CenterService{
		repo: repositories.NewCenterRepository(),
	}
}

func (s *CenterService) Create(db *driver.DB, Center models.Center) (*models.Center, error) {
	err := s.repo.Create(db.Postgres, &Center)
	return &Center, err
}

func (s *CenterService) Fetch(db *driver.DB, params models.QueryParams, example models.Center) (*models.Page, error) {
	return s.repo.FetchWithPagination(db.Postgres, params, example)
}

func (s *CenterService) FetchOne(db *driver.DB, id string) (*models.Center, error) {
	temp, err := s.repo.FetchOne(db.Postgres, id)
	if err != nil {
		return nil, err
	}
	return temp.(*models.Center), err
}
