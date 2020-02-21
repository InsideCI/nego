package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type GenericRepository struct {
	Type interface{}
}

func (r *GenericRepository) output() interface{} {
	out := reflect.New(reflect.TypeOf(r.Type)).Interface()
	return out
}

func (r *GenericRepository) slice() interface{} {
	out := reflect.New(reflect.SliceOf(reflect.TypeOf(r.Type))).Interface()
	return out
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
	v := reflect.ValueOf(model)
	modelType := reflect.TypeOf(model)
	fmt.Println(v)

	modelSlice := reflect.MakeSlice(reflect.SliceOf(modelType), 0, 10)
	fmt.Println(modelSlice.CanAddr())

	if err := db.Limit(limit).Find(&modelSlice).Error; err != nil {
		return nil, err
	}
	return &model, nil
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
