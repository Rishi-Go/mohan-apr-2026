package dto

import (
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type CommentRequest struct {
	Comment string    `json:"comment" gorm:"size:500"`
	BlogID  uuid.UUID `json:"blog_id"`
	UserID  uuid.UUID `json:"user_id"`
}

type CommentResponse struct {
	Comments   []models.Comment `json:"comments"`
	Pagination Pagination       `json:"pagination"`
}

type CommentInsertResponse struct {
	Comments models.Comment `json:"comments"`
}

type CommentInsertResponses struct {
	Comments []models.Comment `json:"comments"`
}
