package repositories

import (
	"fmt"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
)

type GenericRepository struct {
}

func (r *GenericRepository) Create(db *gorm.DB, value interface{}) (interface{}, error) {
	if err := db.Create(&value).Error; err != nil {
		return nil, err
	}
	return value, nil
}

func (r *GenericRepository) Count(db *gorm.DB, value interface{}) (int, error) {
	var count int
	if err := db.Model(value).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *GenericRepository) Fetch(db *gorm.DB, limit int, model interface{}) (interface{}, error) {
	var err error
	if limit != 0 {
		fmt.Print("something to hold on")
		if err = db.Limit(limit).Find(model).Error; err != nil {
			return nil, err
		}
	} else if err = db.Limit(constants.MAXIMUM_FETCH).Find(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *GenericRepository) FetchOne(db *gorm.DB, id int, value interface{}) (interface{}, error) {
	err := db.Where("id = ?", id).First(&value).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r *GenericRepository) Exists(db *gorm.DB, id int, value interface{}) (bool, error) {
	var count int
	if err := db.Model(value).Count(&count).Error; err != nil {
		return false, err
	}
	exists := count != 0
	return exists, nil
}
