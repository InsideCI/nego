package model

// Course abstracts
type Course struct {
	ID          int    `json:"id" gorm:"PRIMARY_KEY;index:index_course_id"`
	Nome        string `json:"nome"`
	Local       string `json:"local"`
	Tipo        string `json:"tipo"`
	Coordenador string `json:"coordenador"`
	IDCentro    int    `json:"idCentro" gorm:"index:index_course_center"`
}
