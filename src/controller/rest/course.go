package rest

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/services"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"net/http"
)

type CourseController struct {
	db      *driver.DB
	service *services.CourseService
}

func NewCourseController(db *driver.DB) *CourseController {
	return &CourseController{
		db:      db,
		service: services.NewCourseService(),
	}
}

func (c *CourseController) Create(w http.ResponseWriter, r *http.Request) {
	course := r.Context().Value("payload").(*models.Course)

	created, err := c.service.Create(c.db, *course)
	if err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *CourseController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(models.QueryParams)
	example := r.Context().Value("example").(*models.Course)

	page, err := c.service.Fetch(c.db, params, *example)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)

}

func (c *CourseController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	student, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
		return
	}

	_ = json.NewEncoder(w).Encode(student)

}
