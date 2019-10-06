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

func (p *postgresCenterRepository) Create(m *model.Center) (int64, error) {

	return 0, nil
}

func (p *postgresCenterRepository) Fetch(num int64) ([]*model.Center, error) {
	return nil, nil
}
