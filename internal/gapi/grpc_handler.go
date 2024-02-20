package gapi

import (
	"sharecycle/internal/usecase"
	"sharecycle/pkg/logger"
)

type GrpcHandler struct {
	Log  logger.Logger
	User usecase.UserInputPort
}

func NewGrpcHandler(
	log logger.Logger,
	user usecase.UserInputPort,
) *GrpcHandler {
	return &GrpcHandler{
		Log:  log,
		User: user,
	}
}
