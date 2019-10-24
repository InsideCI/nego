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

func (r *postgresStudentRespository) Create(student *model.Student) (magtricula int, err error) {
	if err = r.db.Create(&student).Error; err != nil {
		return
	}
	return student.Matricula, nil
}

func (r *postgresStudentRespository) Fetch(limit ...int) (students []*model.Student, err error) {
	if err = r.db.Find(&students).Error; err != nil {
		return
	}
	return
}

func (r *postgresStudentRespository) FetchOne(registration int64) (student *model.Student, err error) {
	var st model.Student
	if err = r.db.First(&st, 11409558).Error; err != nil {
		return
	}
	return
}
