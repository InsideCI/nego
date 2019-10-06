package model

import "github.com/jinzhu/gorm"

// Course abstracts
type Course struct {
	gorm.Model
	ID          int    `json:"courseID"`
	Nome        string `json:"nome"`
	Cidade      string `json:"cidade"`
	Tipo        string `json:"tipo"`
	Coordenador string `json:"coordenador"`
}
