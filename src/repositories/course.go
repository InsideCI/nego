package repositories

import (
	"github.com/InsideCI/nego/src/models"
)

type CourseRepository struct {
	GenericRepository
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{
		struct{ Type interface{} }{Type: models.Course{}},
	}
}
