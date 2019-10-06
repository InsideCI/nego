package handler

import (
	"fmt"
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

func (c *Center) Create(w http.ResponseWriter, r *http.Request) {

	center := model.Center{
		ID:   1321,
		Nome: "eae comparças",
	}

	c.repo.Create(&center)
	fmt.Fprint(w, "User created.")
}
