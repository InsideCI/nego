package repository

import "github.com/InsideCI/nego/src/model"

// StudentRepository abstract basic CRUD methods for students.
type StudentRepository interface {
	Create(students []model.Student) (int, error)
	Fetch(limit int) ([]*model.Student, error)
	FetchOne(registration int) (*model.Student, error)
}
