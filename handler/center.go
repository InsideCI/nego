package handler

import (
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/center"
)

// Center is an abstraction handler for Center type
type Center struct {
	repo repository.CenterRepository
}

// NewCenterHandler returns a Handler of type Center
func NewCenterHandler(driver *driver.DB) *Center {
	return &Center{
		repo: center.NewCenterRepository(driver.psql),
	}
}
