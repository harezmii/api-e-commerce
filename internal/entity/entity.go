package entity

type Faq struct {
	Question string `json:"question" form:"question" validate:"required"`
	Answer   string `json:"answer" form:"answer" validate:"required"`
	Status   *bool  `json:"status" form:"status" validate:"required"`
}

type User struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Surname  string `json:"surname" form:"surname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   *bool  `json:"status" form:"status" validate:"required"`
	//Profile  Profile `json:"profile" form:"profile" validate:"required"`
}

type Profile struct {
	UserId  int    `json:"userId" validate:"required"`
	Image   string `json:"image" form:"image" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
}

type Message struct {
	Name    string `json:"name" form:"name" validate:"required"`
	Email   string `json:"email" form:"email" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
	Subject string `json:"subject" form:"subject" validate:"required"`
	Message string `json:"message" form:"message" validate:"required"`
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
