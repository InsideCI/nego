package repositories

import (
	"github.com/InsideCI/nego/src/models"
	"github.com/jinzhu/gorm"
)

type StudentRepository struct {
	GenericRepository
}

// NewStudentRepository creates a PostgreSQL CRUD interface implementation
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) FetchByName(db *gorm.DB, name string) (*models.Student, error) {
	var student models.Student
	err := db.Where("nome = ?", name).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) FetchByRegistration(db *gorm.DB, registration string) (*models.Student, error) {
	var student models.Student
	err := db.Where("matricula = ?", registration).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
