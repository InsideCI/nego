package department

import (
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/jinzhu/gorm"
)

type postgresDepartmentRepository struct {
	db *gorm.DB
}

// NewDepartmentRepository returns an interface abstraction for CRUD methods.
func NewDepartmentRepository(db *gorm.DB) repository.DepartmentRepository {
	return &postgresDepartmentRepository{
		db: db,
	}
}

func (p *postgresDepartmentRepository) Create(d *model.Department) (int, error) {
	return 0, nil
}

func (p *postgresDepartmentRepository) Fetch(limit int) ([]*model.Department, error) {
	return nil, nil
}
