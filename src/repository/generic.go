package repository

import (
	"github.com/InsideCI/nego/src/model"
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

func (r *GenericRepository) Create(db *gorm.DB, model interface{}) error {
	return db.Create(model).Error
}

func (r *GenericRepository) Count(db *gorm.DB) (int, error) {
	var count int
	out := r.output()
	if err := db.Model(out).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *GenericRepository) Fetch(db *gorm.DB, limit int) (interface{}, error) {
	out := r.slice()
	if err := db.Limit(limit).Find(out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *GenericRepository) FetchWithPagination(db *gorm.DB, params map[string][]string) (*model.Page, error) {
	//TODO: create pagination
	return nil, nil
}

func (r *GenericRepository) FetchOne(db *gorm.DB, id string) (interface{}, error) {
	out := r.output()
	err := db.Where("id = ?", id).First(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *GenericRepository) Exists(db *gorm.DB, model interface{}) bool {
	var count int
	db.Model(model).Count(&count)
	return count != 0
}
