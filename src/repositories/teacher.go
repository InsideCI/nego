package repositories

import "github.com/InsideCI/nego/src/models"

//TeacherRespository represents a repository for Teacher model.
type TeacherRepository struct {
	GenericRepository
}

//NewTeacherRepository returns a new instance of teacher repository.
func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{
		struct{ Type interface{} }{Type: models.Teacher{}},
	}
}
