package model

import "github.com/go-playground/validator/v10"

// Center abstracts a UFPB Center.
type Center struct {
	ID   int    `validate:"required" json:"id" gorm:"PRIMARY_KEY;index:index_center_id"`
	Nome string `validate:"required" json:"nome"`
}

// NewCenter creates a new instance of Center
func NewCenter(id int, nome string) *Center {
	return &Center{
		ID:   id,
		Nome: nome,
	}
}

func (s *Center) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
