package app

import (
	"sharecycle/configs"
	"sharecycle/foundation/database"
	v1 "sharecycle/internal/controller/http/v1"
	"sharecycle/internal/repository"
	"sharecycle/internal/usecase"
	"sharecycle/pkg/logger"
)

func Ready() *Server {

	conf := configs.GetConfig()
	db := database.NewDB(conf)
	l := logger.NewArLogger(conf)

	userRepository := repository.NewUserRepository(db)
	userInputPort := usecase.NewUser(userRepository)

	handler := v1.NewHandler(l, userInputPort)
	sv := NewServer(conf, db, l, handler)

	return sv
}
