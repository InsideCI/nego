package model

import "github.com/go-playground/validator/v10"

type Class struct {
	ID        int    `validate:"required" json:"id" ID   gorm:"PRIMARY_KEY;index:index_center_id"`
	Nome      string `validate:"required" json:"nome"`
	Turma     string `validate:"required" json:"turma"`
	Professor string `validate:"required" json:"professor"`
	Horario   string `validate:"required" json:"horario"`
	IDCurso   int    `validate:"required" json:"idCurso" gorm:"PRIMARY_KEY,index:index_class_course"`
}

func NewClass(id, idCurso int, nome, turma, professor, horario string) *Class {
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
