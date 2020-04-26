package repositories

import (
	"github.com/InsideCI/nego/src/models"
)

type CourseRepository struct {
	*GenericRepository
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{
		GenericRepository: NewGenericRepository(models.Course{}),
	}
}
