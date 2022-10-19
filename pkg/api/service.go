package api

import (
	"30/pkg/model"
	"30/pkg/storage/interfaces"
)

// DTO для объекта запроса
type RequestDTO struct {
	Source int `json:"source_id"`
	Target int `json:"target_id"`
	NewAge int `json:"new age"`
}

// DTO для объекта ответа
type ResponseDTO struct {
	Message string        `json:"message"`
	Items   []*model.User `json:"items"`
}

// DTO для объекта ответа с ошибкой
type ResponseErrorDTO struct {
	Message string `json:"message"`
}

type Server struct {
	Repo *interfaces.Repository
}
