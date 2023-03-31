package dto

type CreateAccountDTO struct {
	Username string `json:"title" form:"title" binding:"required"`
	Email    string `json:"description" form:"description" binding:"required"`
	Password uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type LoginAccountDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password uint64 `json:"password" form:"email" binding:"required"`
}

type UpdateAccountDTO struct {
}
