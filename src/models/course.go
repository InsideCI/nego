package models

import "github.com/go-playground/validator/v10"

// Course abstracts
type Course struct {
	ID          int    `validate:"required" json:"id" gorm:"PRIMARY_KEY;index:index_course_id"`
	Nome        string `validate:"required" json:"nome"`
	Local       string `validate:"required" json:"local"`
	Tipo        string `validate:"required" json:"tipo"`
	Coordenador string `validate:"required" json:"coordenador"`
	IDCentro    int    `validate:"required" json:"idCentro" gorm:"index:index_course_center"`
}

func NewCourse(id, idCentro int, nome, local, tipo, coordenador string) *Course {
	return &Course{
		ID:          id,
		Nome:        nome,
		Local:       local,
		Tipo:        tipo,
		Coordenador: coordenador,
		IDCentro:    idCentro,
	}
}

func (s *Course) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
