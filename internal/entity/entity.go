package entity

type Category struct {
	Title       string `json:"title" xml:"title" form:"title" validate:"required,min=3,max=30"`
	Keywords    string `json:"keywords" xml:"keywords" form:"keywords" validate:"required,min=3,max=250"`
	Description string `json:"description" xml:"description" form:"description" validate:"required,min=10,max=250"`
	Image       string `json:"image" xml:"image" form:"image" validate:"required"`
	Status      *bool  `json:"status" xml:"status" form:"status" validate:"required"`
}

type Faq struct {
	Question string `json:"question" form:"question" validate:"required,min=3,max=40"`
	Answer   string `json:"answer" form:"answer" validate:"required,min=3,max=250"`
	Status   *bool  `json:"status" form:"status" validate:"required"`
}

type User struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Surname  string `json:"surname" form:"surname" validate:"required,min=3"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Status   *bool  `json:"status" form:"status" validate:"required"`
	//Profile  Profile `json:"profile" form:"profile" validate:"required"`
}

type Profile struct {
	UserId  int    `json:"userId,omitempty"`
	Image   string `json:"image" xml:"image,omitempty" form:"image,omitempy"`
	Url     string `json:"url,omitempty" xml:"url,omitempty"`
	Address string `json:"address" form:"address" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required,min=10,max=15"`
}
type Image struct {
	ContentId int    `json:"contentId,omitempty" xml:"contentId"`
	Title     string `json:"title" xml:"title" form:"title" validate:"required,min=3,max=50"`
	Image     string `json:"image" xml:"image" form:"image"`
	Url       string `json:"url,omitempty" xml:"url,omitempty"`
}
type Comment struct {
	Comment   string  `json:"comment" xml:"comment" form:"comment" validate:"required"`
	Rate      float64 `json:"rate" xml:"rate" form:"rate" validate:"required"`
	ContentId int     `json:"contentId,omitempty"`
	UserId    int     `json:"userId,omitempty"`
	IP        string  `json:"ip" xml:"ip" form:"ip" validate:"required"`
	Status    *bool   `json:"status" xml:"status" form:"status" validate:"required"`
}
type Message struct {
	Name    string `json:"name" form:"name" validate:"required,min=3,max=30"`
	Email   string `json:"email" form:"email" validate:"required,email"`
	Phone   string `json:"phone" form:"phone" validate:"required,min=10,max=15"`
	Subject string `json:"subject" form:"subject" validate:"required,min=3,max=20"`
	Message string `json:"message" form:"message" validate:"required,min=5,max=250"`
	IP      string `json:"ip" form:"ip" validate:"required"`
	Status  *bool  `json:"status" form:"status" validate:"required"`
}

type Product struct {
	Title       string `json:"title" xml:"title" form:"title" validate:"required"`
	Keywords    string `json:"keywords" xml:"keywords" form:"keywords" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" validate:"required"`
	Image       string `json:"image,omitempty" xml:"image,omitempty" form:"image"`
	Url         string `json:"url" xml:"url" form:"url" `
	Status      *bool  `json:"status" xml:"status" form:"status" validate:"required"`
}

type Login struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// KAFKA ENTITY

type Kafka struct {
	Topic  string
	Config map[string]interface{}
}
type LogToKafka struct {
	LogLevel   string
	LogMessage string
}
