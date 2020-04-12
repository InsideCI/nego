package models

import "testing"

func TestNewDepartment(t *testing.T) {
	got := NewDepartment(1010, "Department", 1010)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
