package record

import "time"

type AccountLoginHistoriesRecord struct {
	AccountLoginHistoryID uint      `gorm:"primaryKey"`
	AccountID             int       `gorm:"not null"`
	UserID                int       `gorm:"not null"`
	Username              string    `gorm:"default:null;"`
	Password              string    `gorm:"default:null;"`
	Token                 string    `gorm:"default:null;"`
	TokenExpireAt         time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05;"`
	LoginAt               time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05;"`
	LoginOutAt            time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05;"`
	CreatedAt             time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05;"`
	UpdatedAt             time.Time `gorm:"type:datetime;default:null;format:2006-01-02 15:04:05;"`
}

func (AccountLoginHistoriesRecord) TableName() string {
	return "account_login_histories"
}
