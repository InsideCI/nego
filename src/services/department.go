package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type DepartmentService struct {
	repo *repositories.DepartmentRepository
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{
		repo: repositories.NewDepartmentRepository(),
	}
}

func (s *DepartmentService) Create(db *driver.DB, department models.Department) (*models.Department, error) {
	err := s.repo.Create(db.Postgres, &department)
	return &department, err
}

func (s *DepartmentService) Fetch(db *driver.DB, params models.QueryParams, example models.Department) (*models.Page, error) {
	return s.repo.FetchWithPagination(db.Postgres, params, example)
}

func (s *DepartmentService) FetchOne(db *driver.DB, id string) (*models.Department, error) {
	temp, err := s.repo.FetchOne(db.Postgres, id)
	return temp.(*models.Department), err
}
