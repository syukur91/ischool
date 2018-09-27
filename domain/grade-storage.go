package domain

type GradeStorage interface {
	CreateGrade(s *Grade) (*Grade, error)
	GetGrade(id string) (*Grade, error)
	GetGradeByNIS(id string) (*Grade, error)
	UpdateGrade(nis string, s *Grade) (*Grade, error)
}
