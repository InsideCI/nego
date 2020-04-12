package models

import "testing"

func TestNewCenter(t *testing.T) {
	got := NewCenter(1010, "Center")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
