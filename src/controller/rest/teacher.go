package rest

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/services"
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

func (s *TeacherController) Create(w http.ResponseWriter, r *http.Request) {
	teacher := r.Context().Value("payload").(model.Teacher)

	created, err := s.service.Create(s.db, teacher)

	if err != nil {
		log.Println("[HANDLER]", err)
	}

	_ = json.NewEncoder(w).Encode(created)
}
