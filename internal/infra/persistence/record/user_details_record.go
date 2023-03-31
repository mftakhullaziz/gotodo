package record

import "time"

type UserDetailRecord struct {
	UserID      uint      `gorm:"primaryKey"`
	Username    string    `gorm:"not null"`
	Password    string    `gorm:"not null"`
	Email       string    `gorm:"not null"`
	Name        string    `gorm:"not null"`
	MobilePhone int       `gorm:"not null"`
	Address     string    `gorm:"not null"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

func (UserDetailRecord) TableName() string {
	return "user_details"
}
