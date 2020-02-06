package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
	"strconv"
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

	created, err := s.repo.Create(db.Postgres, &student)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *StudentService) Fetch(db *driver.DB, params map[string][]string) (*[]models.Student, error) {
	var fetched *[]models.Student
	var err error

	if name, ok := params["name"]; ok {
		fetched, err = s.repo.FetchByName(db.Postgres, name[0])
	} else {
		if limit, err := strconv.Atoi(params["limit"][0]); err != nil {
			return nil, err
		} else {
			fetched, err = s.repo.Fetch(db.Postgres, limit)
		}
	}

	if err != nil {
		return nil, err
	}
	return fetched, nil
}

func (s *StudentService) FetchByName(db *driver.DB, name string) (*[]models.Student, error) {
	fetched, err := s.repo.FetchByName(db.Postgres, name)
	if err != nil {
		return nil, err
	}
	return fetched, nil
}

func (s *StudentService) FetchOne(db *driver.DB, registration int) (*models.Student, error) {
	fetched, err := s.repo.FetchOne(db.Postgres, registration)
	if err != nil {
		return nil, err
	}
	return fetched, nil
}
