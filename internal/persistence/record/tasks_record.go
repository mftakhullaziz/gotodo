package record

import (
	"time"
)

type TaskRecord struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Completed   bool      `gorm:"not null"`
	CompletedAt time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

func (TaskRecord) TableName() string {
	return "tasks"
}
