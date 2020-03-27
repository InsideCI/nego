package router

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/go-chi/chi"
)

func InitRoutes(db *driver.DB, router *chi.Mux) {

	router.Route("/classes", NewClassRouter(db))
	router.Route("/courses", NewCourseRouter(db))
	router.Route("/departments", NewDepartmentRouter(db))
	router.Route("/students", NewStudentRouter(db))
	router.Route("/teachers", NewTeacherRouter(db))
	router.Route("/centers", NewCenterRouter(db))

}
