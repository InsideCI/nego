package repository

import (
	"github.com/InsideCI/nego/src/model"
	"github.com/jinzhu/gorm"
)

type ClassRepository struct {
}

func NewClassRepository() *ClassRepository {
	return &ClassRepository{}
}

func (r *ClassRepository) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&model.Class{}).Count(&count).Error
	return count, err
}
