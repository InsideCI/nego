package models_test

import (
	"github.com/InsideCI/nego/src/models"
	"testing"
)

func TestNewTeacher(t *testing.T) {
	got := models.NewTeacher(1010, 1010, "Teacher", "Doctor")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
