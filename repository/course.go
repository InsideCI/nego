package repository

import "github.com/InsideCI/nego/model"

// CourseRepository
type CourseRepository interface {
	Create(c *model.Course) (int, error)
	Fetch(limit int) ([]*model.Course, error)
}
