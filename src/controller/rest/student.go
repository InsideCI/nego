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

// StudentController is a controller that wraps a StudentService.
type StudentController struct {
	db      *driver.DB
	service *services.StudentService
}

// NewStudentController creates a new controller with a specific database engine.
func NewStudentController(db *driver.DB) *StudentController {
	return &StudentController{
		db:      db,
		service: services.NewStudentService(),
	}
}

func (c *StudentController) Create(w http.ResponseWriter, r *http.Request) {
	student := r.Context().Value("payload").(*models.Student)

	created, err := c.service.Create(c.db, *student)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *StudentController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(models.QueryParams)
	example := r.Context().Value("example").(*models.Student)

	page, err := c.service.Fetch(c.db, params, *example)
	if err != nil {
		utils.Throw(w, exceptions.InternalError, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)
}

// FetchOne uses id param as primary key search
func (c *StudentController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	student, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.NotFound, err)
		return
	}
	_ = json.NewEncoder(w).Encode(student)
}
