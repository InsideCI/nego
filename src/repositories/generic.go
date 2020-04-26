package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories/cache"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
)

//GenericRepository abstracts all basic crud methods.
type GenericRepository struct {
	Type    interface{}
	caching *cache.BadgerRepository
}

//NewGenericRepository returns a new instance of a generic repository with caching.
func NewGenericRepository(t interface{}) *GenericRepository {
	return &GenericRepository{
		Type:    t,
		caching: cache.NewBadgerRepository(t),
	}
}

func (r *GenericRepository) output() interface{} {
	out := reflect.New(reflect.TypeOf(r.Type)).Interface()
	return out
}

func (r *GenericRepository) slice() interface{} {
	out := reflect.New(reflect.SliceOf(reflect.TypeOf(r.Type))).Interface()
	return out
}

//Create inserts a new instance of model Type on the database.
func (r *GenericRepository) Create(db *gorm.DB, model interface{}) error {
	return db.Create(model).Error
}

//Count returns the total ammount of registers of model Type.
func (r *GenericRepository) Count(db *gorm.DB) (int, error) {
	var count int
	out := r.output()
	if err := db.Model(out).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//Fetch returns all registered models Type from database.
func (r *GenericRepository) Fetch(db *gorm.DB) (interface{}, error) {
	out := r.slice()
	if err := db.Find(out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

//FetchWithPagination returns a page of Type based on query parameters.
func (r *GenericRepository) FetchWithPagination(db *gorm.DB, params models.QueryParams, example interface{}) (*models.Page, error) {
	var (
		out         = r.slice()
		offset      int
		limit       int
		payloadSize int
		fields, _   = model.Fields(example)
	)

	if params.Limit <= 0 || params.Limit > constants.MaximumFetch {
		limit = constants.MaximumFetch
	} else {
		limit = params.Limit
	}

	if params.Page < 0 {
		offset = 0
	} else {
		offset = params.Page * limit
	}

	tx := db.Where("")

	// Parse filter parameters.
	if !model.IsZero(example) {
		for _, field := range fields {
			if value, _ := model.Get(example, field.Name); value != "" && value != 0 {
				if reflect.TypeOf(value).Kind() != reflect.String {
					tx = tx.Where(strcase.ToSnake(field.Tag.Get("json"))+" = ?", value)
				} else {
					valueStr := fmt.Sprintf("%v", value)
					// Case insensitive and unrelative filter position on register field.
					tx = tx.Where(strcase.ToSnake(field.Tag.Get("json"))+" ILIKE ?", "%"+valueStr+"%")
				}
			}
		}
	}

	counted := make(chan bool, 1)
	go func() {
		tx.Model(out).Count(&payloadSize)
		counted <- true
	}()

	// Parse sort parameters.
	if len(params.Order) != 0 {
		for _, field := range params.Order {
			for _, exampleField := range fields {
				if field == exampleField.Tag.Get("json") {
					tx = tx.Order(strcase.ToSnake(field))
				}
			}
		}
	}

	tx = tx.Offset(offset).Limit(limit).Find(out)
	<-counted

	totalPages := (payloadSize / limit) + 1

	return models.NewPage(payloadSize, limit, params.Page, totalPages, reflect.ValueOf(out).Elem().Len(), out), nil
}

//FetchOne returns a instance of a model Type by it's ID.
func (r *GenericRepository) FetchOne(db *gorm.DB, id string) (interface{}, error) {
	out := r.output()

	cache, err := r.caching.Get(id)

	if err == nil {
		json.Unmarshal(cache, out)
		return out, nil
	}

	fmt.Println(cache)

	err = db.Where("id = ?", id).First(out).Error
	if err != nil {
		return nil, err
	}
	if model.IsZero(out) {
		return nil, errors.New("register not found")
	}

	err = r.caching.SaveByIDKey(id, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

//Exists checks if a register exits by it's ID.
func (r *GenericRepository) Exists(db *gorm.DB, id string) bool {
	out := r.output()
	var count int
	db.Model(out).Where("id = ?", id).Count(&count)
	return count != 0
}
