package domain

type StudentStorage interface {
	CreateStudent(s *Student) (*Student, error)
}
