package models

import "testing"

func TestNewClass(t *testing.T) {
	got := NewClass(1010, "1010", "Name", "Class", "Teacher", "Time")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
