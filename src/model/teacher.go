package model

import "github.com/go-playground/validator/v10"

type Teacher struct {
	ID             int    `validate:"required" json:"id" gorm:"PRIMARY_KEY;type:bigint;"`
	Nome           string `validate:"required" json:"nome" gorm:"index:index_teacher_name"`
	Grau           string `json:"grau"`
	IDDepartamento int    `validate:"required" json:"grau"`
}

func NewTeacher(id, idDepartamento int, nome, grau string) *Teacher {
	return &Teacher{
		ID:             id,
		Nome:           nome,
		Grau:           grau,
		IDDepartamento: idDepartamento,
	}
}

func (s *Teacher) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
