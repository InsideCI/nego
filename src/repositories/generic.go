package repositories

import (
	"github.com/InsideCI/nego/src/models"
	"github.com/jinzhu/gorm"
)

type GenericRepository struct {
}

func (r *GenericRepository) Create(db *gorm.DB, value interface{}) (created int, err error) {
	if err = db.Create(&value).Error; err != nil {
		return
	}
	return
}

func (r *GenericRepository) Count(db *gorm.DB, value interface{}) int {
	//return db.Count(value)
}

func (r *GenericRepository) Fetch(db *gorm.DB, limit ...int) (students []*models.Student, err error) {
	//TODO: implement default maximum API limit if not provided
	if err = db.Find(&students).Error; err != nil {
		return
	}
	return
}

func (r *GenericRepository) FetchOne(db *gorm.DB, id int, value *interface{}) (interface{}, error) {
	err := db.Where("id = ?", id).First(&value).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}
