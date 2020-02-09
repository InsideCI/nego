package repository

import (
	"github.com/InsideCI/nego/src/model"
	"github.com/jinzhu/gorm"
)

type CourseRepository struct {
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

func (r *CourseRepository) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&model.Course{}).Count(&count).Error
	return count, err
}
