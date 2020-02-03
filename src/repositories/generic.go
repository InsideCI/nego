package repositories

import (
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
)

type GenericRepository struct {
}

func (r *GenericRepository) Create(db *gorm.DB, value interface{}) interface{} {
	//if err = db.Create(&value).Error; err != nil {
	//	return nil, err
	//}
	//return value, nil
	db.Create(&value)
	return value
}

func (r *GenericRepository) Count(db *gorm.DB, value interface{}) (int, error) {
	var count int
	if err := db.Model(value).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *GenericRepository) Fetch(db *gorm.DB, limit int, value interface{}) (interface{}, error) {
	var err error
	if limit != 0 {
		if err = db.Limit(limit).Find(&value).Error; err != nil {
			return nil, err
		}
	} else if err = db.Limit(constants.MAXIMUM_FETCH).Find(&value).Error; err != nil {
		return nil, err
	}
	return value, nil
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
