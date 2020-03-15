package repositories

import "github.com/InsideCI/nego/src/models"

type TeacherRepository struct {
	GenericRepository
}

func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{
		struct{ Type interface{} }{Type: models.Teacher{}},
	}
}
