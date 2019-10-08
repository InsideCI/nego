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

func NewCourseRepository(db *driver.DB) repository.CourseRepository {
	return &postgresCourseRepository{
		db: db.Psql,
	}
}

func (r *postgresCourseRepository) Create(course *model.Course) (id int, err error) {
	if err = r.db.Create(&course).Error; err != nil {
		return
	}
	return course.ID, nil
}

func (r *postgresCourseRepository) Fetch(limit int) (courses []*model.Course, err error) {
	if err = r.db.Find(&courses).Error; err != nil {
		return
	}
	return
}
