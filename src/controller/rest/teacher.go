package rest

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/services"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"log"
	"net/http"
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
	teacher := r.Context().Value("payload").(model.Teacher)

	created, err := c.service.Create(c.db, teacher)

	if err != nil {
		log.Println("[HANDLER]", err)
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *TeacherController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	student, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
		return
	}

	_ = json.NewEncoder(w).Encode(student)

}
