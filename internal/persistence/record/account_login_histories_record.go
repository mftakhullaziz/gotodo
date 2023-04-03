package record

import "time"

type AccountLoginHistoriesRecord struct {
	AccountLoginHistoryID uint      `gorm:"primaryKey"`
	AccountID             int       `gorm:"not null"`
	UserID                int       `gorm:"not null"`
	Username              string    `gorm:"default:null;"`
	Password              string    `gorm:"default:null;"`
	Token                 string    `gorm:"default:null;"`
	LoginAt               time.Time `gorm:"default:null;"`
	LoginOutAt            time.Time `gorm:"default:null;"`
	CreatedAt             time.Time `gorm:"default:null;"`
	UpdatedAt             time.Time `gorm:"default:null;"`
}

func (AccountLoginHistoriesRecord) TableName() string {
	return "account_login_histories"
}
