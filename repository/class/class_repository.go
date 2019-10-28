package class

import (
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/jinzhu/gorm"
	"log"
)

type postgresClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *driver.DB) repository.ClassRepository {
	return &postgresClassRepository{
		db: db.Psql,
	}
}

func (r *postgresClassRepository) Create(classes []model.Class) (created int, err error) {
	for _, class := range classes {
		if err = r.db.Create(&class).Error; err != nil {
			return
		}
		created++
	}
	return
}

func (r *postgresClassRepository) Fetch(limit int) (classes []model.Class, err error) {
	if err := r.db.Find(&classes).Error; err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *postgresClassRepository) FetchOne(id int) (class model.Class, err error) {
	if err := r.db.Where("id = ?", id).First(&class).Error; err != nil {
		return
	}
	return
}

func (r *postgresClassRepository) FetchByCourse(courseID int) (classes []model.Class, err error) {
	if err := r.db.Where("id_curso = ?", courseID).Find(&classes).Error; err != nil {
		return
	}
	return
}
