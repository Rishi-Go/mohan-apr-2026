package dto

import (
	"blog_post/pkg/models"

	"github.com/gofrs/uuid"
)

type LikeRequest struct {
	LikeResponse string    `json:"like_response" validate:"required,oneof= Yes No"`
	UserID       uuid.UUID `json:"user_id"`
}

type LikeResponse struct {
	Like       []models.Like `json:"like"`
	Pagination Pagination    `json:"pagination"`
}
