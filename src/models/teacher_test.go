package models

import "testing"

func TestNewTeacher(t *testing.T) {
	got := NewTeacher(1010, 1010, "Teacher", "Doctor")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
