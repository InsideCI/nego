package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

// Student abstracts a basic UFPB student.
type Student struct {
	Matricula int    `validate:"required" json:"matricula" gorm:"PRIMARY_KEY;type:bigint;index:index_student_id;"`
	Nome      string `validate:"required" json:"nome" gorm:"index:index_student_name"`
	IDCurso   int    `validate:"required" json:"idCurso" gorm:"index:index_student_course"`
}

// NewStudent creates a new instance of a student for test purposes.
func NewStudent(matricula int, nome string, idCurso int) *Student {
	return &Student{
		Matricula: matricula,
		Nome:      nome,
		IDCurso:   idCurso,
	}
}

// JSON encodes a Student model as a JSON.
func (s *Student) JSON() []byte {
	p, _ := json.Marshal(s)
	return p
}

func (s *Student) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
