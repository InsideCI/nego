package repositories

import (
	"github.com/InsideCI/nego/src/models"
)

//StudentRepository abstracts a basic Student repository.
type StudentRepository struct {
	*GenericRepository
}

// NewStudentRepository creates a PostgreSQL CRUD interface implementation
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		GenericRepository: NewGenericRepository(models.Student{}),
	}
}
