package repository

import (
	"github.com/InsideCI/nego/model"
)

// CenterRepository abstract database CRUD methods for Center.
type CenterRepository interface {
	Create(centers []model.Center) (int, error)
	Fetch(limit int) ([]*model.Center, error)
}
