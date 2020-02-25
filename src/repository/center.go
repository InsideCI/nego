package repository

import (
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
)

type CenterRepository struct {
	GenericRepository
}

func NewCenterRepository() *CenterRepository {
	return &CenterRepository{
		struct{ Type interface{} }{Type: model.Center{}},
	}
}

func (r *CenterRepository) FetchByName(db *gorm.DB, name string) (*[]model.Center, error) {
	var centers []model.Center
	err := db.Limit(constants.LIMIT_FAST_FETCH).Where("nome LIKE ?", "%"+name+"%").Order("nome").Find(&centers).Error
	if err != nil {
		return nil, err
	}
	return &centers, nil
}
