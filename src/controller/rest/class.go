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

type ClassController struct {
	db      *driver.DB
	service *services.ClassService
}

func NewClassController(db *driver.DB) *ClassController {
	return &ClassController{
		db:      db,
		service: services.NewClassService(),
	}
}

func (c *ClassController) Create(w http.ResponseWriter, r *http.Request) {
	class := r.Context().Value("payload").(*models.Class)

	created, err := c.service.Create(c.db, *class)
	if err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *ClassController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(models.QueryParams)
	example := r.Context().Value("example").(*models.Class)

	page, err := c.service.Fetch(c.db, params, *example)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)

}

func (c *ClassController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	student, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
		return
	}

	_ = json.NewEncoder(w).Encode(student)

}
