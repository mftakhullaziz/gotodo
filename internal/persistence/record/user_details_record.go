package record

import "time"

type UserDetailRecord struct {
	UserID      uint      `gorm:"primaryKey"`
	Username    string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Name        string    `gorm:"not null"`
	MobilePhone int       `gorm:"not null"`
	Address     string    `gorm:"not null"`
	Status      string    `gorm:"default:null;"`
	CreatedAt   time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05"`
	UpdatedAt   time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05"`
}

func (UserDetailRecord) TableName() string {
	return "user_details"
}
