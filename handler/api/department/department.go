package department

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/department"
)

// Department abstracts an UFPB Department
type Department struct {
	repo repository.DepartmentRepository
}

// NewDepartmentHandler creates a new handler for Department
func NewDepartmentHandler(driver *driver.DB) *Department {
	return &Department{
		repo: department.NewDepartmentRepository(driver.Psql),
	}
}

// Create receives and request body and creates a new instance of Department.
func (d *Department) Create(w http.ResponseWriter, r *http.Request) {
	var deps []model.Department
	if err := json.NewDecoder(r.Body).Decode(&deps); err != nil {
		log.Println(err)
		return
	}
	d.repo.Create(deps)
	w.Write([]byte("OK"))
}

// Fetch returns an array containing departments
func (d *Department) Fetch(w http.ResponseWriter, r *http.Request) {
	deps, err := d.repo.Fetch(10)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(deps)
}
