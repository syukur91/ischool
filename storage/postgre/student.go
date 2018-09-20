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

func (s *StudentPotgreStorage) GetStudent(id string) (*domain.Student, error) {

	t := domain.Student{}
	d := domain.StudentData{}

	if err := s.DB.Where("id=?", id).Find(&d).Error; err != nil {
		return nil, errors.Errorf("APIKey with id: %s is not found", id)
	}

	if err := json.Unmarshal([]byte(d.Data), &t); err != nil {
		return nil, errors.Errorf("APIKey with id: %s error unmarshall", id)
	}

	return &t, nil
}

func (s *StudentPotgreStorage) GetStudentByNIS(nis string) (*domain.Student, error) {

	t := domain.Student{}
	d := domain.StudentData{}

	if err := s.DB.Where(`data@>'{"nis": "` + nis + `"}'`).Find(&d).Error; err != nil {
		return nil, errors.Errorf("Student with nis: %s is not found", nis)
	}

	if err := json.Unmarshal([]byte(d.Data), &t); err != nil {
		return nil, errors.Errorf("Student with nis: %s error unmarshall", nis)
	}

	return &t, nil
}

func (s *StudentPotgreStorage) UpdateStudent(nis string, d *domain.Student) (*domain.Student, error) {

	b, err := json.Marshal(d)
	if err != nil {
		return nil, errors.Errorf("Error marshalling struct id %s", d.ID)
	}

	value := string(b)

	update := s.DB.Model(&domain.StudentData{}).Where(`data@>'{"nis": "`+nis+`"}'`).Update("data", value)
	if update.Error != nil {
		return nil, errors.Errorf(update.Error.Error(), nis)
	}

	return d, nil
}
