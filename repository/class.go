package repository

import "github.com/InsideCI/nego/model"

type ClassRepository interface {
	Create(classes []model.Class) (created int, err error)
	Fetch(limit int) (class []model.Class, err error)
	FetchOne(id int) (class model.Class, err error)
	FetchByCourse(courseID int) (class []model.Class, err error)
}
