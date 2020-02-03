package repositories

type CourseRepository struct {
	GenericRepository
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}
