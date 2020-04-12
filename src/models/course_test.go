package models

import "testing"

func TestNewCourse(t *testing.T) {
	got := NewCourse(1010, 1010, "Course", "Place", "Type", "Coordinator")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
