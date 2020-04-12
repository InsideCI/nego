package models

import "testing"

func TestNewPage(t *testing.T) {
	got := NewPage(1, 1, 1, 1, 1, " ")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
