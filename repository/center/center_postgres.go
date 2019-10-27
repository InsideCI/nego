package center

import (
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/jinzhu/gorm"
)

type postgresCenterRepository struct {
	db *gorm.DB
}

// NewCenterRepository returns implementation of Center store interface.
func NewCenterRepository(db *gorm.DB) repository.CenterRepository {
	return &postgresCenterRepository{
		db: db,
	}
}

func (p *postgresCenterRepository) Create(centers []model.Center) (created int, err error) {
	for _, center := range centers {
		if err = p.db.Create(&center).Error; err != nil {
			return
		}
		created++
	}
	return
}

func (p *postgresCenterRepository) Fetch(num int) (centers []*model.Center, err error) {
	if err = p.db.Find(&centers).Error; err != nil {
		return
	}
	return
}
