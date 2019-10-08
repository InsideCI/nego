package repository

import "github.com/InsideCI/nego/model"

// DepartmentRepository abstracts CRUD methods for Department.
type DepartmentRepository interface {
	Create(d *model.Department) (int, error)
	Fetch(limit int) ([]*model.Department, error)
}
