package model_test

import (
	"github.com/InsideCI/nego/src/model"
	"testing"
)

func TestNewClass(t *testing.T) {
	got := model.NewClass(1010, 1010, "Name", "Class", "Teacher", "Time")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
