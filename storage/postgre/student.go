package student

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/syukur91/ischool/domain"
)

type StudentPotgreStorage struct {
	Log *log.Entry
	DB  *gorm.DB
}

func NewStudentPostgreStorage(db *gorm.DB, log *log.Entry) *StudentPotgreStorage {

	client := &StudentPotgreStorage{
		Log: log,
		DB:  db,
	}
	return client
}

func (s *StudentPotgreStorage) CreateStudent(d *domain.Student) (*domain.Student, error) {

	b, err := json.Marshal(d)
	if err != nil {
		return nil, errors.Errorf("Error marshalling struct id %s", d.ID)
	}

	rec := &domain.StudentData{ID: d.ID, Data: string(b), CreatedAt: d.CreatedAt}
	create := s.DB.Create(rec)
	if create.Error != nil {
		return nil, create.Error
	}

	return d, nil
}
