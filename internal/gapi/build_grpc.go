package gapi

import (
	"sharecycle/configs"
	"sharecycle/foundation/database"
	"sharecycle/internal/repository"
	"sharecycle/internal/usecase"
	"sharecycle/pkg/logger"
)

func GrpcReady() *Server {

	conf := configs.GetConfig()
	db := database.NewDB(conf)
	l := logger.NewArLogger(conf)

	userRepository := repository.NewUserRepository(db)
	userInputPort := usecase.NewUser(userRepository)

	handler := NewGrpcHandler(l, userInputPort)
	sv := NewGrpcServer(conf, db, l, handler)

	return sv
}
