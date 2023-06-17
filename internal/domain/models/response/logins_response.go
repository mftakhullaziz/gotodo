package response

import "time"

type LoginResponse struct {
	AccountID         int       `json:"account_id"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	LoginCreationTime time.Time `json:"login_at"`
	LoginToken        string    `json:"token"`
}
