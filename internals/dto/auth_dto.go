package dto

import "blog_post/pkg/models"

type SignUpRequest struct {
	FirstName string `json:"firstname" gorm:"size:100"`
	LastName  string `json:"lastname" gorm:"size:100"`
	UserName  string `json:"username" gorm:"size:100"`
	Password  string `json:"password" gorm:"size:100"` // bcrypt
	Email     string `json:"email"`
}

type UserResponse struct {
	BlogUsers  []models.BlogUsers `json:"blog_users"`
	Pagination Pagination         `json:"pagination"`
}
