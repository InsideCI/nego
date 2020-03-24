package repositories

import (
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/jinzhu/gorm"
	"strings"
)

type CourseRepository struct {
	GenericRepository
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{
		struct{ Type interface{} }{Type: models.Course{}},
	}
}

func (r *CourseRepository) FetchByName(db *gorm.DB, name string) (*[]models.Course, error) {
	var courses []models.Course
	name = strings.ToUpper(name)
	err := db.Limit(constants.LIMIT_FAST_FETCH).Where("nome LIKE ?", "%"+name+"%").Order("nome").Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return &courses, nil
}