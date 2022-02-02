package entity

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
	UserId  int    `json:"userId" validate:"required"`
	Image   string `json:"image" form:"image" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required,min=10,max=15"`
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
