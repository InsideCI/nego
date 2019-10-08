package repository

import (
	"github.com/InsideCI/nego/model"
)

// CenterRepository abstract database CRUD methods for Center.
type CenterRepository interface {
	Create(p *model.Center) (int, error)
	Fetch(limit int) ([]*model.Center, error)
}
