package router

import (
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/utils"
	"github.com/go-chi/chi"
)

var auth = utils.NewJWT()

//InitRoutes instantiates all available Nego routes.
func InitRoutes(db *driver.DB, router *chi.Mux) {

	router.Use(auth.Verifier())

	router.Route("/", NewAuthRouter(db))
	router.Route("/classes", NewClassRouter(db))
	router.Route("/courses", NewCourseRouter(db))
	router.Route("/departments", NewDepartmentRouter(db))
	router.Route("/students", NewStudentRouter(db))
	router.Route("/teachers", NewTeacherRouter(db))
	router.Route("/centers", NewCenterRouter(db))

}
