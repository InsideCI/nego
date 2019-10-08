package handler

import (
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/course"
	"net/http"
)

type Course struct {
	repo repository.CourseRepository
}

func NewCourseHandler(db *driver.DB) *Course {
	return &Course{
		repo: course.NewCourseRepository(db),
	}
}

func (c *Course) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *Course) Fetch(w http.ResponseWriter, r *http.Request) {

}
