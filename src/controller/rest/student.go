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

// StudentController is a controller that wraps a StudentRepository.
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
	student := r.Context().Value("payload").(model.Student)

	created, err := c.service.Create(c.db, student)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	_ = json.NewEncoder(w).Encode(created)
}

func (c *StudentController) Fetch(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(map[string][]string)

	page, err := c.service.Fetch(c.db, params)
	if err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}
	_ = json.NewEncoder(w).Encode(page)
}

// FetchByMatricula uses id param as primary key search
func (c *StudentController) FetchOne(w http.ResponseWriter, r *http.Request) {
	registration := r.Context().Value("id").(string)

	student, err := c.service.FetchOne(c.db, registration)
	if err != nil {
		utils.Throw(w, exceptions.InvalidAttributes, err)
		return
	}

	_ = json.NewEncoder(w).Encode(student)

}
