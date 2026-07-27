package dto

import (
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type BlogRequest struct {
	Title      string    `json:"title" gorm:"size:100"`
	Content    string    `json:"content" gorm:"size:500"`
	CategoryId uuid.UUID `json:"category_id"`
	AuthorID   uuid.UUID `json:"author_id"`
}

type BlogResponse struct {
	Blog       []models.Blog `json:"blog"`
	Pagination Pagination    `json:"pagination"`
}
