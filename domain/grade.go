package domain

import (
	"time"

	"github.com/google/uuid"
)

// Grade is
type Grade struct {
	ID         string     `json:"id,omitempty"`
	StudentID  string     `json:"student_id,omitempty"`
	GradeValue int        `json:"grade_value,omitempty"`
	CreatedAt  *time.Time `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	UpdatedAt  *time.Time `json:"updated_at"`
	UpdatedBy  string     `json:"updated_by"`
	ExpiredAt  *time.Time `json:"expired_at"`
}

type GradeData struct {
	ID        string
	StudentID string
	Data      string
	CreatedAt *time.Time
}

func (g *Grade) SetCreated(user string) {
	id := uuid.Must(uuid.NewRandom())
	g.ID = id.String()
	now := time.Now()
	g.CreatedAt = &now
	g.CreatedBy = user
	g.UpdatedAt = nil
	g.UpdatedBy = ""
}

func (g *Grade) SetUpdated(user string) {
	now := time.Now()
	g.UpdatedAt = &now
	g.UpdatedBy = user
}
