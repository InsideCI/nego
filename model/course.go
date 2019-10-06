package model

// Course abstracts
type Course struct {
	ID          int    `json:"courseID"`
	Nome        string `json:"nome"`
	Cidade      string `json:"cidade"`
	Tipo        string `json:"tipo"`
	Coordenador string `json:"coordenador"`
}
