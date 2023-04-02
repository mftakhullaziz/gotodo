package dto

import "time"

type AccountDTO struct {
	AccountID uint      `json:"account_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
