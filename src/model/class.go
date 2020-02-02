package model

type Class struct {
	ID        int    `json:"id" ID   gorm:"PRIMARY_KEY;index:index_center_id"`
	Nome      string `json:"nome"`
	Turma     string `json:"turma"`
	Professor string `json:"professor"`
	Horario   string `json:"horario"`
	IDCurso   int    `json:"id_curso" gorm:"PRIMARY_KEY,index:index_class_course"`
}

func NewClass(id, idCurso int, nome, turma, professor, horario string) *Class {
	return &Class{
		ID:        id,
		Nome:      nome,
		Turma:     turma,
		Professor: professor,
		Horario:   horario,
		IDCurso:   idCurso,
	}
}
