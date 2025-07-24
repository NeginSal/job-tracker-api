package job

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Description string
	Company     string
	UserID      string `gorm:"type:uuid;not null"` // foreign key
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
