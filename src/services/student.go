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
	err := s.repo.Create(db.Postgres, &student)
	return &student, err
}

func (s *StudentService) Fetch(db *driver.DB, params models.QueryParams, example models.Student) (*models.Page, error) {
	return s.repo.FetchWithPagination(db.Postgres, params, example)
}

func (s *StudentService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}

func (s *StudentService) FetchOne(db *driver.DB, id string) (*models.Student, error) {
	tmp, err := s.repo.FetchOne(db.Postgres, id)
	if err != nil {
		return nil, err
	}
	return tmp.(*models.Student), err
}

func (s *StudentService) Exists(db *driver.DB, student *models.Student) bool {
	return s.repo.Exists(db.Postgres, student)
}
