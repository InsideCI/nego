package student

import (
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/jinzhu/gorm"
)

type postgresStudentRespository struct {
	db *gorm.DB
}

// NewStudentRepository creates a PostgreSQL CRUD interface implementation
func NewStudentRepository(db *driver.DB) repository.StudentRepository {
	return &postgresStudentRespository{
		db: db.Psql,
	}
}

func (r *postgresStudentRespository) Create(students []model.Student) (created int, err error) {
	for _, student := range students {
		if err = r.db.Create(&student).Error; err != nil {
			return
		}
		created++
	}
	return
}

func (r *postgresStudentRespository) Fetch(limit ...int) (students []*model.Student, err error) {
	if err = r.db.Find(&students).Error; err != nil {
		return
	}
	return
}

func (r *postgresStudentRespository) FetchOne(registration int64) (*model.Student, error) {
	var student model.Student
	err := r.db.Where("matricula = ?", registration).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
