package repositories

import (
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
	"strings"
)

type StudentRepository struct {
}

// NewStudentRepository creates a PostgreSQL CRUD interface implementation
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
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

func (r *StudentRepository) FetchByRegistration(db *gorm.DB, registration string) (*models.Student, error) {
	var student models.Student
	err := db.Where("matricula = ?", registration).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Create(db *gorm.DB, student *models.Student) (*models.Student, error) {
	if err := db.Create(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepository) Count(db *gorm.DB, student models.Student) (int, error) {
	var count int
	if err := db.Model(student).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *StudentRepository) Fetch(db *gorm.DB, limit int) (*[]models.Student, error) {
	var students []models.Student
	if err := db.Limit(limit).Order("nome").Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepository) FetchOne(db *gorm.DB, id int) (*models.Student, error) {
	var student models.Student
	err := db.Where("matricula = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Exists(db *gorm.DB, id int, student models.Student) (bool, error) {
	var count int
	if err := db.Model(student).Count(&count).Error; err != nil {
		return false, err
	}
	exists := count != 0
	return exists, nil
}
