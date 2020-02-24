package repository

import (
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
	"strings"
)

type StudentRepository struct {
	GenericRepository
}

// NewStudentRepository creates a PostgreSQL CRUD interface implementation
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		struct{ Type interface{} }{Type: model.Student{}},
	}
}

func (r *StudentRepository) FetchByName(db *gorm.DB, name string) (*[]model.Student, error) {
	var students []model.Student
	name = strings.ToUpper(name)
	err := db.Limit(constants.LIMIT_FAST_FETCH).Where("nome LIKE ?", "%"+name+"%").Order("nome").Find(&students).Error
	if err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepository) FetchByRegistration(db *gorm.DB, registration string) (*model.Student, error) {
	var student model.Student
	err := db.Where("matricula = ?", registration).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) FetchByIdCourse(db *gorm.DB, idCourse string) ([]model.Student, error) {
	var students []model.Student
	err := db.Where("idCurso = ?", idCourse).First(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
