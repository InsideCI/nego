package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type CourseService struct {
	repo *repositories.CourseRepository
}

func NewCourseService() *CourseService {
	return &CourseService{
		repo: repositories.NewCourseRepository(),
	}
}

func (s *CourseService) Create(db *driver.DB, course models.Course) (*models.Course, error) {
	//TODO: check if relations exists
	err := s.repo.Create(db.Postgres, &course)
	return &course, err
}

func (s *CourseService) Fetch(db *driver.DB, params models.QueryParams, example models.Student) (*models.Page, error) {

	return s.repo.FetchWithPagination(db.Postgres, params, example)

}

func (s *CourseService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}
