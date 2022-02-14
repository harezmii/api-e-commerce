package dto

type LoginUserDTO struct {
	UserID  int    `json:"id" form:"id" validate:"required"`
	Name    string `json:"name" form:"name" validate:"required"`
	Surname string `json:"surname" form:"surname" validate:"required"`
	Email   string `json:"email" form:"email" validate:"required"`
	Status  *bool  `json:"status" form:"status" validate:"required"`
}

type FaqDto struct {
	Id       int    `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
	Status   *bool  `json:"status,omitempty"`
}
type UserDto struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Email   string `json:"email,omitempty"`
	Status  *bool  `json:"status,omitempty"`
}
type CategoryDto struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Status      *bool  `json:"status,omitempty"`
}

type CommentDto struct {
	Id      int     `json:"id,omitempty"`
	Comment string  `json:"comment,omitempty"`
	Rate    float32 `json:"rate,omitempty"`
	Ip      string  `json:"ip,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}
type MessageDto struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Subject string `json:"subject,omitempty"`
	Message string `json:"message,omitempty"`
	IP      string `json:"ip,omitempty"`
	Status  *bool  `json:"status,omitempty"`
}
type ProfileDto struct {
	Id      int    `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Image   string `json:"image,omitempty"`
}
