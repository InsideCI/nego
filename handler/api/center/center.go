package center

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/model"
	"github.com/InsideCI/nego/repository"
	"github.com/InsideCI/nego/repository/center"
)

// Center is an abstraction handler for Center type
type Center struct {
	repo repository.CenterRepository
}

// NewCenterHandler returns a Handler of type Center
func NewCenterHandler(driver *driver.DB) *Center {
	return &Center{
		repo: center.NewCenterRepository(driver.Psql),
	}
}

// Create receives a body composed of an Center JSON data.s
func (c *Center) Create(w http.ResponseWriter, r *http.Request) {
	var centers []model.Center
	if err := json.NewDecoder(r.Body).Decode(&centers); err != nil {
		log.Println("[HANDLER]", err)
		return
	}
	c.repo.Create(centers)
	w.Write([]byte("OK"))
}

// Fetch returns an array containing exactly the quantity of predefined entities.
func (c *Center) Fetch(w http.ResponseWriter, r *http.Request) {
	centers, err := c.repo.Fetch(10)
	if err != nil {
		log.Fatal("[HANDLER]", err)
	}
	json.NewEncoder(w).Encode(centers)
}
