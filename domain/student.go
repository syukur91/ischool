package domain

import (
	"time"

	"github.com/google/uuid"
)

// StudentData is
type Student struct {
	ID        string     `json:"id,omitempty"`
	NIS       string     `json:"nis,omitempty"`
	Name      string     `json:"name,omitempty"`
	Age       int        `json:"age,omitempty"`
	Class     string     `json:"class,omitempty"`
	Sex       string     `json:"sex,omitempty"`
	Grade     int        `json:"grade,omitempty"`
	Address   string     `json:"address,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type StudentData struct {
	ID        string
	Data      string
	CreatedAt *time.Time
}

func (c *Student) SetCreated(user string) {
	id := uuid.Must(uuid.NewRandom())
	c.ID = id.String()
	now := time.Now()
	c.CreatedAt = &now
	c.CreatedBy = user
	c.UpdatedAt = nil
	c.UpdatedBy = ""
}

func (c *Student) SetUpdated(user string) {
	now := time.Now()
	c.UpdatedAt = &now
	c.UpdatedBy = user
}
