package rest

import (
	"encoding/json"
	"net/http"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/services"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/exceptions"
)

type TeacherController struct {
	db      *driver.DB
	service *services.TeacherService
}

func NewTeacherController(db *driver.DB) *TeacherController {
	return &TeacherController{
		db:      db,
		service: services.NewTeacherService(),
	}
}

func (c *TeacherController) Create(w http.ResponseWriter, r *http.Request) {
	teacher := r.Context().Value("payload").(*models.Teacher)

	created, err := c.service.Create(c.db, *teacher)
	if err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *TeacherController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(models.QueryParams)
	example := r.Context().Value("example").(*models.Teacher)

	page, err := c.service.Fetch(c.db, params, *example)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)

}

func (c *TeacherController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	teacher, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.NotFound, err)
		return
	}

	_ = json.NewEncoder(w).Encode(teacher)

}
