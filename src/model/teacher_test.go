package model_test

import (
	"github.com/InsideCI/nego/src/model"
	"testing"
)

func TestNewTeacher(t *testing.T) {
	got := model.NewTeacher(1010, 1010, "Teacher", "Doctor")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
