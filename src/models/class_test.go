package models_test

import (
	"github.com/InsideCI/nego/src/models"
	"testing"
)

func TestNewClass(t *testing.T) {
	got := models.NewClass(1010, "1010", "Name", "Class", "Teacher", "Time")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
