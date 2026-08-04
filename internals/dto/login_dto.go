package dto

type LogInRequest struct {
	UserName string `json:"user_name" gorm:"size:20"`
	Password string `json:"password" gorm:"size:100"`
}
