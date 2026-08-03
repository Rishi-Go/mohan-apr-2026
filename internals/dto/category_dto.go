package dto

import "blog_post/pkg/models"

type CategoryRequest struct {
	CategoryName string `json:"category_name" gorm:"size:100"`
	Description  string `json:"description" gorm:"size:500"`
}

type CategoryResponse struct {
	Category   []models.Category `json:"category"`
	Pagination Pagination        `json:"pagination"`
}
