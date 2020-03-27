package models

import "github.com/go-playground/validator/v10"

type Class struct {
	ID        string `validate:"required" json:"id" gorm:"PRIMARY_KEY;index:index_class_id"`
	Nome      string `validate:"required" json:"nome" gorm:"index:index_class_name"`
	Turma     string `validate:"required" json:"turma"`
	Professor string `json:"professor"`
	Horario   string `json:"horario"`
	IDCurso   int    `validate:"required" json:"idCurso" gorm:"index:index_class_course"`
}

func NewClass(idCurso int, id, nome, turma, professor, horario string) *Class {
	return &Class{
		ID:        id,
		Nome:      nome,
		Turma:     turma,
		Professor: professor,
		Horario:   horario,
		IDCurso:   idCurso,
	}
}

func (s *Class) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
