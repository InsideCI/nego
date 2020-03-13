package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
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

func (s *CourseService) Create(db *driver.DB, course model.Course) (*model.Course, error) {
	//TODO: check if relations exists
	err := s.repo.Create(db.Postgres, &course)
	return &course, err
}

func (s *CourseService) Fetch(db *driver.DB, params map[string][]string) (*model.Page, error) {
	var err error
	var fetched *[]model.Course
	var totalCourses int

	if totalCourses, err = s.Count(db); err != nil {
		return nil, err
	}

	if name, ok := params["name"]; ok {
		fetched, err = s.repo.FetchByName(db.Postgres, name[0])
	} else {
		return s.repo.FetchWithPagination(db.Postgres, params)
	}

	return model.NewPage(totalCourses, len(*fetched), fetched), nil
}

func (s *CourseService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}
