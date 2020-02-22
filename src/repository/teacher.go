package repository

import "github.com/InsideCI/nego/src/model"

type TeacherRepository struct {
	GenericRepository
}

func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{
		struct{ Type interface{} }{Type: model.Teacher{}},
	}
}
