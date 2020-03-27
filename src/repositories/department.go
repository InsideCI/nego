package repositories

import "github.com/InsideCI/nego/src/models"

type DepartmentRepository struct {
	GenericRepository
}

func NewDepartmentRepository() *DepartmentRepository {
	return &DepartmentRepository{GenericRepository{Type: models.Department{}}}
}
