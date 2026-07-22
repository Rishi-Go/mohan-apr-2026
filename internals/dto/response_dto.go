package dto

import "github.com/gofrs/uuid"

type ResponseMessage struct {
	Message string `json:"message"`
}

type Response struct {
	Message string    `json:"message"`
	ID      uuid.UUID `json:"id"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Pagination struct {
	Page   int `json:"page"`
	Limit  int `json:"limit" gorm:"size:99"`
	Total  int `json:"total"`
	Offset int `json:"offset" gorm:"size:99"`
}
