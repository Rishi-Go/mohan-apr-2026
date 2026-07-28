package dto

type LogInRequest struct {
	UserName string `json:"username" gorm:"size:20"`
	Password string `json:"password" gorm:"size:100"`
}
