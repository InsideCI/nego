package models_test

import (
	"github.com/InsideCI/nego/src/models"
	"testing"
)

func TestNewCourse(t *testing.T) {
	got := models.NewCourse(1010, 1010, "Course", "Place", "Type", "Coordinator")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
