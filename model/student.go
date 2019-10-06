package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// Student abstracts a basic UFPB student.
type Student struct {
	gorm.Model
	Matricula string `json:"matricula"`
	Nome      string `json:"nome"`
	CursoID   string `json:"cursoID"`
}

// NewStudent creates a new instance of a student for test purposes.
func NewStudent(matricula string, nome string, cursoID string) *Student {
	return &Student{
		Matricula: matricula,
		Nome:      nome,
		CursoID:   cursoID,
	}
}

// JSON encodes a Student model as a JSON.
func (s *Student) JSON() []byte {
	p, _ := json.Marshal(s)
	return p
}
