package model_test

import (
	"github.com/InsideCI/nego/src/model"
	"testing"
)

func TestNewCenter(t *testing.T) {
	got := model.NewCenter(1010, "Center")
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
