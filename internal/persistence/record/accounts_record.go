package record

import "time"

type AccountRecord struct {
	AccountID uint      `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (AccountRecord) TableName() string {
	return "accounts"
}
