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

func (p *postgresCenterRepository) Create(center *model.Center) (id int, err error) {
	if err = p.db.Create(&center).Error; err != nil {
		return
	}
	return center.ID, nil
}

func (p *postgresCenterRepository) Fetch(num int64) (centers []*model.Center, err error) {
	if err = p.db.Find(&centers).Error; err != nil {
		return
	}
	return
}
