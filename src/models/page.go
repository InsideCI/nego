package models

import "github.com/go-playground/validator/v10"

//Page describes a basic page model for fetching.
type Page struct {
	Total       int         `json:"totalElements"`
	TotalPages  int         `json:"totalPages"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PayloadSize int         `json:"payloadSize"`
	Payload     interface{} `json:"payload"`
}

//NewPage returns a new instance of page.
func NewPage(total, limit, page, totalPages, payloadSize int, payload interface{}) *Page {
	return &Page{
		Total:       total,
		Limit:       limit,
		Page:        page,
		TotalPages:  totalPages,
		PayloadSize: payloadSize,
		Payload:     payload,
	}
}

//Valid checks for fields errors.
func (s *Page) Valid() error {
	v := validator.New()
	return v.Struct(s)
}
