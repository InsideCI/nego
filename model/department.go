package model

// Department abstract department in UFPB
type Department struct {
	ID       int    `json:"id" gorm:"PRIMARY_KEY"`
	Nome     string `json:"nome"`
	IDCentro int    `json:"id_centro"`
}

// NewDepartment returns a new instance of Department
func NewDepartment(id int, nome string, idCentro int) *Department {
	return &Department{
		ID:       id,
		Nome:     nome,
		IDCentro: idCentro,
	}
}
