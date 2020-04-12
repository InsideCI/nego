package models

import "testing"

func TestNewStudent(t *testing.T) {
	got := NewStudent(1010, "Student", 1010)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
