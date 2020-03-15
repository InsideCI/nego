package models_test

import (
	"github.com/InsideCI/nego/src/models"
	"testing"
)

func TestNewCenter(t *testing.T) {
	got := models.NewCenter(1010, "Center")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
