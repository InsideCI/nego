package models

// Center abstracts a UFPB Center.
type Center struct {
	ID   int    `json:"id" gorm:"PRIMARY_KEY;index:index_center_id"`
	Nome string `json:"nome"`
}

// NewCenter creates a new instance of Center
func NewCenter(id int, nome string) *Center {
	return &Center{
		ID:   id,
		Nome: nome,
	}
}