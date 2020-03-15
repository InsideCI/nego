package repositories

import (
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
)

type CenterRepository struct {
	GenericRepository
}

func NewCenterRepository() *CenterRepository {
	return &CenterRepository{
		struct{ Type interface{} }{Type: models.Center{}},
	}
}

func (r *CenterRepository) FetchByName(db *gorm.DB, name string) (*[]models.Center, error) {
	var centers []models.Center
	err := db.Limit(constants.LIMIT_FAST_FETCH).Where("nome LIKE ?", "%"+name+"%").Order("nome").Find(&centers).Error
	if err != nil {
		return nil, err
	}
	return &centers, nil
}
