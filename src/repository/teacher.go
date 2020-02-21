package repository

type TeacherRepository struct {
	GenericRepository
}

func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{}
}
