package model_test

import (
	"github.com/InsideCI/nego/src/model"
	"testing"
)

func TestNewCourse(t *testing.T) {
	got := model.NewCourse(1010, 1010, "Course", "Place", "Type", "Coordinator")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
