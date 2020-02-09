package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService() *CourseService {
	return &CourseService{
		repo: repository.NewCourseRepository(),
	}
}

func (s *CourseService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}
