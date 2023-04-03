package response

import "time"

type LoginResponse struct {
	AccountID uint      `json:"account_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LoginAt   time.Time `json:"login_at"`
	Token     string    `json:"token"`
}
