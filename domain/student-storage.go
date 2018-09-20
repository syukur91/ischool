package domain

type StudentStorage interface {
	CreateStudent(s *Student) (*Student, error)
	GetStudent(id string) (*Student, error)
	GetStudentByNIS(id string) (*Student, error)
	UpdateStudent(nis string, s *Student) (*Student, error)
}
