package repositories

import (
	"fmt"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
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

func (r *GenericRepository) FetchWithPagination(db *gorm.DB, params models.QueryParams, example interface{}) (*models.Page, error) {
	var (
		out         = r.slice()
		offset      int
		limit       int
		payloadSize int
	)

	if params.Limit <= 0 || params.Limit > constants.MAXIMUM_FETCH {
		limit = constants.MAXIMUM_FETCH
	} else {
		limit = params.Limit
	}

	if params.Page < 0 {
		offset = 0
	} else {
		offset = params.Page * limit
	}

	tx := db.Debug().Where("")

	if !model.IsZero(example) {
		fields, _ := model.Fields(example)
		for _, field := range fields {
			if value, _ := model.Get(example, field.Name); value != nil && value != 0 {
				valueStr := fmt.Sprintf("%v", value)
				tx = tx.Where(field.Tag.Get("json")+" ILIKE ?", "%"+valueStr+"%")
			}
		}
	}
	tx.Model(out).Count(&payloadSize)

	//TODO: implement sort by using params.Order

	tx = tx.Offset(offset).Limit(limit)

	if err := tx.Find(out).Error; err != nil {
		return nil, err
	}

	totalPages := payloadSize / limit

	return models.NewPage(payloadSize, limit, params.Page, totalPages, reflect.ValueOf(out).Elem().Len(), out), nil
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