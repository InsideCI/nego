package repositories

import "github.com/InsideCI/nego/src/models"

type ClassRepository struct {
	GenericRepository
}

func NewClassRepository() *ClassRepository {
	return &ClassRepository{
		struct{ Type interface{} }{Type: models.Class{}},
	}
}
