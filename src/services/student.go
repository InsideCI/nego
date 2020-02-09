package services

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/repository"
	"strconv"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService() *StudentService {
	return &StudentService{
		repo: repository.NewStudentRepository(),
	}
}

func (s *StudentService) Create(db *driver.DB, student model.Student) (*model.Student, error) {
	//TODO: check and throw error if student' course doesn't exists in the database.
	//courseRepository := repository.NewCourseRepository()
	return s.repo.Create(db.Postgres, &student)
}

func (s *StudentService) Fetch(db *driver.DB, params map[string][]string) (*model.Page, error) {
	var err error
	var fetched *[]model.Student
	var totalStudents int

	if totalStudents, err = s.Count(db); err != nil {
		return nil, err
	}

	if name, ok := params["name"]; ok {
		fetched, err = s.repo.FetchByName(db.Postgres, name[0])
	} else {
		if limit, err := strconv.Atoi(params["limit"][0]); err != nil {
			return nil, err
		} else {
			fetched, err = s.repo.Fetch(db.Postgres, limit)
		}
	}
	return model.NewPage(totalStudents, len(*fetched), fetched), nil
}

func (s *StudentService) FetchByName(db *driver.DB, name string) (*[]model.Student, error) {
	return s.repo.FetchByName(db.Postgres, name)
}

func (s *StudentService) FetchOne(db *driver.DB, registration string) (*model.Student, error) {
	return s.repo.FetchOne(db.Postgres, registration)
}

func (s *StudentService) Delete(db *driver.DB, registration int) (bool, error) {
	return s.repo.Delete(db.Postgres, registration)
}

func (s *StudentService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres)
}
