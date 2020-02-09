package model_test

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/model"
	"testing"
)

var (
	id      = 1010
	nome    = "Student"
	idCurso = 1010
)

func TestNewStudent(t *testing.T) {
	got := model.NewStudent(id, nome, idCurso)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}

func TestStudent_JSON(t *testing.T) {
	jsonStudent := model.NewStudent(id, nome, idCurso).JSON()

	var student model.Student
	if err := json.Unmarshal(jsonStudent, &student); err != nil {
		t.Error(err)
	}
}
