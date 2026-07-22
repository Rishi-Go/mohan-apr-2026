package dto

import "time"

type CategoryRequest struct {
	CategoryName string    `json:"category_name" gorm:"size:100"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
