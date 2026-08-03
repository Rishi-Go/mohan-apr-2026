package dto

import "blog_post/pkg/models"

type SignUpRequest struct {
	FirstName string `json:"firstname"  validate:"required,min=3"`
	LastName  string `json:"lastname" validate:"required,min=1"`
	UserName  string `json:"username" validate:"required,min=3,max=20" gorm:"unique;not null"`
	Password  string `json:"password" validate:"required,min=6" gorm:"not null"`
	Email     string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Role      string `json:"role"`
}

type UserResponse struct {
	BlogUsers  []models.BlogUsers `json:"blog_users"`
	Pagination Pagination         `json:"pagination"`
}
