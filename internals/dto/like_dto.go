package dto

import "github.com/gofrs/uuid"

type LikeRequest struct {
	LikeResponse bool      `json:"like_response"`
	UserID       uuid.UUID `json:"user_id"`
}
