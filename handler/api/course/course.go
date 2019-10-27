package course

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/course"
)

//Course abstracts a graduation course on UFPB.
type Course struct {
	repo repository.CourseRepository
}

// NewCourseHandler creates a customized handler for Courses.
func NewCourseHandler(db *driver.DB) *Course {
	return &Course{
		repo: course.NewCourseRepository(db),
	}
}

// Create uses request body as parameter for a new Course entry.
func (c *Course) Create(w http.ResponseWriter, r *http.Request) {
	var courses []model.Course
	if err := json.NewDecoder(r.Body).Decode(&courses); err != nil {
		log.Println(err)
		return
	}
	c.repo.Create(courses)
	w.Write([]byte("OK"))
}

// Fetch returns all courses available.
func (c *Course) Fetch(w http.ResponseWriter, r *http.Request) {
	courses, err := c.repo.Fetch(10)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(courses)
}
