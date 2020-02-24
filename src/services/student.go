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
	err := s.repo.Create(db.Postgres, &student)
	return &student, err
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
		limit, err := strconv.Atoi(params["limit"][0])
		temp, err := s.repo.Fetch(db.Postgres, limit)
		if err != nil {
			return nil, err
		}
		fetched = temp.(*[]model.Student)
	}

	return model.NewPage(totalStudents, len(*fetched), fetched), nil
}

func (s *StudentService) Count(db *driver.DB) (int, error) {
	return s.repo.Count(db.Postgres, &model.Student{})
}

func (s *StudentService) FetchOne(db *driver.DB, id string) (*model.Student, error) {
	return s.repo.FetchByRegistration(db.Postgres, id)
}

func (s *StudentService) Exists(db *driver.DB, student *model.Student) bool {
	return s.repo.Exists(db.Postgres, student)
}
