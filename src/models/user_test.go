package models

import "testing"

func TestNewUser(t *testing.T) {
	got := NewUser("name", "lastName", "password", "email", "token")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
