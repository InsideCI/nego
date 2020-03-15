package models_test

import (
	"github.com/InsideCI/nego/src/models"
	"testing"
)

func TestNewStudent(t *testing.T) {
	got := models.NewStudent(1010, "Student", 1010)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
