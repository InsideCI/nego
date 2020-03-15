package models_test

import (
	"github.com/InsideCI/nego/src/models"
	"testing"
)

func TestNewDepartment(t *testing.T) {
	got := models.NewDepartment(1010, "Department", 1010)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
