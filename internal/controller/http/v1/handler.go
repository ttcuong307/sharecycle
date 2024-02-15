package v1

import (
	"sharecycle/internal/usecase"
	"sharecycle/pkg/logger"
)

type Handler struct {
	Log  logger.Logger
	User usecase.UserInputPort
}

func NewHandler(
	log logger.Logger,
	user usecase.UserInputPort,
) *Handler {
	return &Handler{
		Log:  log,
		User: user,
	}
}
