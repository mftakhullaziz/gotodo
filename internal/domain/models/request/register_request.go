package request

type RegisterRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
