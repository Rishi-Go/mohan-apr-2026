package dto

import (
	"blog_post/pkg/models"
)

type SignUpRequest struct {
	FirstName string          `json:"first_name"  validate:"required,min=3" gorm:"not null"`
	LastName  string          `json:"last_name" validate:"required,min=1" gorm:"not null"`
	UserName  string          `json:"user_name" validate:"required,min=3,max=20" gorm:"not null"`
	Password  string          `json:"password" validate:"required,min=6" gorm:"not null"`
	Email     string          `json:"email" validate:"required,email" gorm:"unique;not null"`
	Role      models.UserRole `json:"role"`
}

type UserResponse struct {
	BlogUsers  []models.BlogUsers `json:"blog_users"`
	Pagination Pagination         `json:"pagination"`
}

type BlogUserResponse struct {
	BlogUsers models.BlogUsers `json:"blog_users"`
}
type BlogUsersResponse struct {
	BlogUsers []models.BlogUsers `json:"blog_users"`
}