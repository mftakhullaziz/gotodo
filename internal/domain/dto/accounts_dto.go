package dto

type CreateAccountDTO struct {
	Username string `json:"username" form:"title" binding:"required"`
	Email    string `json:"email" form:"description" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginAccountDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password uint64 `json:"password" form:"email" binding:"required"`
}

type UpdateAccountDTO struct {
}
