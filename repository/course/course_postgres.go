package course

import (
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/jinzhu/gorm"
)

type postgresCourseRepository struct {
	db *gorm.DB
}

// NewCourseRepository creates an instance of CRUD methods for Course
// using postgres.
func NewCourseRepository(db *driver.DB) repository.CourseRepository {
	return &postgresCourseRepository{
		db: db.Postgres,
	}
}

func (r *postgresCourseRepository) Create(courses []model.Course) (created int, err error) {
	for _, course := range courses {
		if err = r.db.Create(&course).Error; err != nil {
			return created, err
		}
		created++
	}
	return
}

func (r *postgresCourseRepository) Fetch(limit int) (courses []*model.Course, err error) {
	if err = r.db.Find(&courses).Error; err != nil {
		return
	}
	return
}
