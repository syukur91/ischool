package student

type StudentData struct {
	Name string `json:"name,omitempty"`
	Age  string `json:"age,omitempty"`
}

type Student interface {
	GetStudents(class string) ([]string, error)
}
