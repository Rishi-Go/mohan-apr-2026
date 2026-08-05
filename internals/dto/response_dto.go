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

type ErrorMessage struct {
	Message    string    `json:"message"`
	StatusCode int       `json:"status_code"`
	ID         uuid.UUID `json:"ID"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit" gorm:"size:99"`
	Total int `json:"total"`
}

type TokenMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Token      string `json:"token"`
}

type SuccessResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data "`
}
