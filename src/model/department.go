package model

import "github.com/go-playground/validator/v10"

// Department abstract department in UFPB
type Department struct {
	ID       int    `validate:"required" json:"id" gorm:"PRIMARY_KEY;index:index_department_id"`
	Nome     string `validate:"required" json:"nome"`
	IDCentro int    `validate:"required" json:"idCentro" gorm:"index:index_dept_center"`
}

// NewDepartment returns a new instance of Department
func NewDepartment(id int, nome string, idCentro int) *Department {
	return &Department{
		ID:       id,
		Nome:     nome,
		IDCentro: idCentro,
	}
}

func (s *Department) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
