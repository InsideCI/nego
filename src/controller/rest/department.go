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

type DepartmentController struct {
	db      *driver.DB
	service *services.DepartmentService
}

func NewDepartmentController(db *driver.DB) *DepartmentController {
	return &DepartmentController{
		db:      db,
		service: services.NewDepartmentService(),
	}
}

func (c *DepartmentController) Create(w http.ResponseWriter, r *http.Request) {
	department := r.Context().Value("payload").(*models.Department)

	created, err := c.service.Create(c.db, *department)
	if err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *DepartmentController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(models.QueryParams)
	example := r.Context().Value("example").(*models.Department)

	page, err := c.service.Fetch(c.db, params, *example)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)

}

func (c *DepartmentController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	department, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
		return
	}

	_ = json.NewEncoder(w).Encode(department)

}
