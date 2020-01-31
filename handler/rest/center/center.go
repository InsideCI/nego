package center

import (
	"encoding/json"
	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/center"
	"log"
	"net/http"
)

// Center is an abstraction handler for Center type
type Center struct {
	repo repository.CenterRepository
}

// NewCenterHandler returns a Handler of type Center
func NewCenterHandler(driver *driver.DB) *Center {
	return &Center{
		repo: center.NewCenterRepository(driver.Postgres),
	}
}

// Create receives a body composed of an Center JSON data.s
func (c *Center) Create(w http.ResponseWriter, r *http.Request) {
	var centers []model.Center
	if err := json.NewDecoder(r.Body).Decode(&centers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_, err := c.repo.Create(centers)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	w.Write([]byte("OK"))
}

// Fetch returns an array containing exactly the quantity of predefined entities.
func (c *Center) Fetch(w http.ResponseWriter, r *http.Request) {
	centers, err := c.repo.Fetch(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}
	json.NewEncoder(w).Encode(centers)
}

// FetchOne returns an Center by it's ID.
func (c *Center) FetchOne(w http.ResponseWriter, r *http.Request) {

}
