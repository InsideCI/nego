package repositories

import (
	"github.com/InsideCI/nego/src/models"
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
		struct{ Type interface{} }{Type: models.Student{}},
	}
}

func (r *StudentRepository) FetchByName(db *gorm.DB, name string) (*[]models.Student, error) {
	var students []models.Student
	name = strings.ToUpper(name)
	err := db.Limit(constants.LIMIT_FAST_FETCH).Where("nome LIKE ?", "%"+name+"%").Order("nome").Find(&students).Error
	if err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepository) FetchByIdCourse(db *gorm.DB, idCourse string) ([]models.Student, error) {
	var students []models.Student
	err := db.Where("idCurso = ?", idCourse).First(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
