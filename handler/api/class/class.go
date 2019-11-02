package class

import (
	"encoding/json"
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/class"
	"log"
	"net/http"
	"strconv"
)

type Class struct {
	repo repository.ClassRepository
}

// NewClassHandler returns a new Handler for a predefined database ORM.
func NewClassHandler(driver *driver.DB) *Class {
	return &Class{
		repo: class.NewClassRepository(driver.Psql),
	}
}

func (c *Class) Create(w http.ResponseWriter, r *http.Request) {
	var classes []model.Class
	if err := json.NewDecoder(r.Body).Decode(&classes); err != nil {
		log.Println("[HANDLER]", err)
	}
	created, err := c.repo.Create(classes)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	w.Write([]byte(strconv.Itoa(created)))
}
