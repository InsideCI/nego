package models

// Student abstracts a basic member of a UFPB student
type Student struct {
	matricula string
	name      string
	cursoID   string
}

// NewStudent creates a new instance of a student
func (s *Student) NewStudent(matricula string, name string, cursoID string) *Student {
	return &Student{
		matricula: matricula,
		name:      name,
		cursoID:   cursoID,
	}
}
