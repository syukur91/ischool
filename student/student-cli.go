package student

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type StudentCLI struct {
	Log *log.Entry
	DB  *gorm.DB
}

func NewStudentCLI(log *log.Entry) *StudentCLI {

	client := &StudentCLI{
		Log: log,
	}

	return client
}

func (s *StudentCLI) GetStudents(class string) ([]string, error) {

	var students []string

	students = append(students, "Dilan")
	students = append(students, "Milea")

	return students, nil
}
