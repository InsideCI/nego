package repository

import (
	"github.com/InsideCI/nego/src/model"
	"github.com/jinzhu/gorm"
	"reflect"
	"strconv"
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
	var (
		out    = r.slice()
		total  int
		offset int
		err    error
	)

	if total, err = r.Count(db); err != nil {
		return nil, err
	}

	limit, _ := strconv.Atoi(params["limit"][0])
	if limit < 0 {
		limit = 0
	}
	page, _ := strconv.Atoi(params["page"][0])
	if page < 0 {
		offset = 0
	} else {
		offset = page * limit
	}

	if err := db.Offset(offset).Limit(limit).Find(out).Error; err != nil {
		return nil, err
	}

	totalPages := total / limit

	return model.NewPage(total, offset, page, totalPages, out), nil
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
