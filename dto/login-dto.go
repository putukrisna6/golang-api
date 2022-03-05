package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"requried" validate:"min:6"`
}
