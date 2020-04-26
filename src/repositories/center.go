package repositories

import (
	"github.com/InsideCI/nego/src/models"
)

type CenterRepository struct {
	*GenericRepository
}

func NewCenterRepository() *CenterRepository {
	return &CenterRepository{
		GenericRepository: NewGenericRepository(models.Center{}),
	}
}
