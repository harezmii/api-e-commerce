package dto

type LoginUserDTO struct {
	UserID  int    `json:"id" form:"id" validate:"required"`
	Name    string `json:"name" form:"name" validate:"required"`
	Surname string `json:"surname" form:"surname" validate:"required"`
	Email   string `json:"email" form:"email" validate:"required"`
	Status  *bool  `json:"status" form:"status" validate:"required"`
}

type FaqDto struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Status   *bool  `json:"status"`
}
type UserDto struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Status  *bool  `json:"status"`
}
type MessageDto struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	IP      string `json:"ip"`
	Status  *bool  `json:"status"`
}
type ProfileDto struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Image   string `json:"image"`
}
