package model

// Course abstracts
type Course struct {
	ID          int    `json:"id" gorm:"index:index_course_id"`
	Nome        string `json:"nome"`
	Cidade      string `json:"cidade"`
	Tipo        string `json:"tipo"`
	Coordenador string `json:"coordenador"`
	IDCentro    int    `json:"id_centro" gorm:"index:index_course_center"`
}
