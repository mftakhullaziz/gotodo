package record

import "time"

type AccountLoginHistoriesRecord struct {
	AccountLoginHistoryID uint      `gorm:"primaryKey"`
	AccountID             int       `gorm:"not null"`
	UserID                int       `gorm:"not null"`
	Email                 string    `gorm:"null"`
	Password              string    `gorm:"null"`
	LoginStatus           string    `gorm:"null"`
	LoginAt               time.Time `gorm:"null"`
	LoginOutAt            time.Time `gorm:"null"`
	CreatedAt             time.Time `gorm:"null"`
	UpdatedAt             time.Time `gorm:"null"`
}

func (AccountLoginHistoriesRecord) TableName() string {
	return "account_login_histories"
}
