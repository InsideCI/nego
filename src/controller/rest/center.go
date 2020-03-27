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

type CenterController struct {
	db      *driver.DB
	service *services.CenterService
}

func NewCenterController(db *driver.DB) *CenterController {
	return &CenterController{
		db:      db,
		service: services.NewCenterService(),
	}
}

func (c *CenterController) Create(w http.ResponseWriter, r *http.Request) {
	center := r.Context().Value("payload").(*models.Center)

	created, err := c.service.Create(c.db, *center)
	if err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *CenterController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(models.QueryParams)
	example := r.Context().Value("example").(*models.Center)

	page, err := c.service.Fetch(c.db, params, *example)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)

}

func (c *CenterController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	center, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
		return
	}

	_ = json.NewEncoder(w).Encode(center)

}
