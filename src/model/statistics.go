package model

import "github.com/go-playground/validator/v10"

type GeneralStatistic struct {
	Students    int
	Courses     int
	Departments int
	Centers     int
	Teachers    int
}

func (g *GeneralStatistic) Valid() error {
	v := validator.New()
	return v.Struct(g)
}

func NewGeneralStatistic(students, courses, departments, centers, Teachers int) *GeneralStatistic {
	return &GeneralStatistic{
		Students:    students,
		Courses:     courses,
		Departments: departments,
		Centers:     centers,
		Teachers:    Teachers,
	}
}
