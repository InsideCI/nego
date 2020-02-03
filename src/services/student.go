package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService() *StudentService {
	return &StudentService{
		repo: repositories.NewStudentRepository(),
	}
}

func (s *StudentService) Create(db *driver.DB, student models.Student) (*models.Student, error) {
	//TODO: check and throw error if student' course doesn't exists in the database.
	//courseRepository := repositories.NewCourseRepository()
	var created models.Student
	var err error
	//TODO: enable context pass through methods
	created = s.repo.Create(db.Postgres, student).(models.Student)

	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (s *StudentService) FetchByName(db *driver.DB) ([]models.Student, error) {

}
