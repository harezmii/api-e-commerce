package dto

type LoginUserDTO struct {
	UserID  int    `json:"id" form:"id" validate:"required"`
	Name    string `json:"name" form:"name" validate:"required"`
	Surname string `json:"surname" form:"surname" validate:"required"`
	Email   string `json:"email" form:"email" validate:"required"`
	Status  *bool  `json:"status" form:"status" validate:"required"`
}
