package services

import (
	"github.com/InsideCI/nego/src/repositories"
)

type CourseService struct {
	repo *repositories.GenericRepository
}

func NewCourseService() *CourseService {
	return &CourseService{}
}
