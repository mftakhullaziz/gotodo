package record

import (
	"time"
)

type TaskRecord struct {
	TaskID      uint      `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Completed   bool      `gorm:"not null"`     // Flagging if completed true task is completed and false task in progress
	TaskStatus  string    `gorm:"default:null"` // Flagging to task inactive or active
	CompletedAt time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05"`
	CreatedAt   time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05"`
	UpdatedAt   time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05"`
}

func (TaskRecord) TableName() string {
	return "tasks"
}
