package dto

import (
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type ReplyRequest struct {
	Reply     string    `json:"reply" gorm:"size:500"`
	CommentID uuid.UUID `json:"comment_id"`
	UserID    uuid.UUID `json:"user_id"`
}

type ReplyResponse struct {
	Reply      []models.Reply `json:"reply"`
	Pagination Pagination     `json:"pagination"`
}

type ReplyInsertResponse struct {
	Reply models.Reply `json:"reply"`
}

type ReplyInsertResponses struct {
	Reply []models.Reply `json:"reply"`
}
