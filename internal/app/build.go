package app

import (
	"sharecycle/configs"
	"sharecycle/foundation/database"
)

func Ready() *Server {

	conf := configs.GetConfig()
	db := database.NewDB(conf)
	sv := NewServer(conf, db)

	return sv
}
