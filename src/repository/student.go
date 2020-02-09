package repository

import (
	"github.com/InsideCI/nego/src/model"
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

func (r *StudentRepository) Create(db *gorm.DB, student *model.Student) (*model.Student, error) {
	if err := db.Create(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepository) Count(db *gorm.DB) (int, error) {
	var count int
	if err := db.Model(&model.Student{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *StudentRepository) Fetch(db *gorm.DB, limit int) (*[]model.Student, error) {
	var students []model.Student
	if err := db.Limit(limit).Order("nome").Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepository) FetchOne(db *gorm.DB, registration string) (*model.Student, error) {
	var student model.Student
	err := db.Where("matricula = ?", registration).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Exists(db *gorm.DB, id int) (bool, error) {
	var count int
	if err := db.Model(&model.Student{Matricula: id}).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}

func (r *StudentRepository) Delete(db *gorm.DB, id int) (bool, error) {
	if err := db.Delete(&model.Student{Matricula: id}).Error; err != nil {
		return false, err
	}
	return true, nil
}
