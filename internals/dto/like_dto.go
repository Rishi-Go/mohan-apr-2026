package dto

import (
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type LikeRequest struct {
	LikeResponse bool      `json:"like_response"`
	BlogID       uuid.UUID `json:"blog_id"`
	UserID       uuid.UUID `json:"user_id"`
}

type LikeResponse struct {
	Like       []models.Like `json:"like"`
	Pagination Pagination    `json:"pagination"`
}

type LikeInsertResponse struct {
	Like models.Like `json:"like"`
}

type LikeInsertResponses struct {
	Like []models.Like `json:"like"`
}
