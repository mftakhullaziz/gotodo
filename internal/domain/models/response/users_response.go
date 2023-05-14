package response

type UserDetailResponse struct {
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	MobilePhone int    `json:"mobile_phone"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
