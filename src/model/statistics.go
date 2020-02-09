package model

type GeneralStatistic struct {
	Students    int
	Courses     int
	Departments int
	Centers     int
	Teachers    int
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
