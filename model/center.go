package model

// Center abstracts a UFPB Center.
type Center struct {
	ID   int    `json:"id" gorm:"PRIMARY_KEY"`
	Nome string `json:"nome"`
}

// NewCenter creates a new instance of Center
func NewCenter(id int, nome string) *Center {
	return &Center{
		ID:   id,
		Nome: nome,
	}
}
