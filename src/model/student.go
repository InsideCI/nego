package model

import (
	"github.com/go-playground/validator/v10"
)

// Student abstracts a basic UFPB student.
type Student struct {
	ID      int    `validate:"required" json:"id" gorm:"PRIMARY_KEY;type:bigint;"`
	Nome    string `validate:"required" json:"nome" gorm:"index:index_student_name"`
	IDCurso int    `validate:"required" json:"idCurso" gorm:"index:index_student_course"`
}

// NewStudent creates a new instance of a student for test purposes.
func NewStudent(matricula int, nome string, idCurso int) *Student {
	return &Student{
		ID:      matricula,
		Nome:    nome,
		IDCurso: idCurso,
	}
}

func (s *Student) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
