package repository

import "github.com/InsideCI/nego/model"

// StudentRepository abstract basic CRUD methods for students.
type StudentRepository interface {
	Create(s *model.Student) (int, error)
	Fetch(limit ...int) ([]*model.Student, error)
	FetchOne(registration int64) (*model.Student, error)
}
