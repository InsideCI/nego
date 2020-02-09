package model_test

import (
	"github.com/InsideCI/nego/src/model"
	"testing"
)

func TestNewDepartment(t *testing.T) {
	got := model.NewDepartment(1010, "Department", 1010)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
