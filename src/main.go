package main

import (
	"log"
	"net"
	"time"

	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/ezio1119/fishapp-chat/infrastructure"
	"github.com/ezio1119/fishapp-chat/infrastructure/middleware"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
)

func main() {
	dbConn := infrastructure.NewGormConn()
	r := infrastructure.NewRedisClient()
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	c := controllers.NewChatController(
		interactor.NewChatInteractor(
			dbConn,
			r,
			time.Duration(conf.C.Sv.Timeout)*time.Second,
		),
	)
	server := infrastructure.NewGrpcServer(middleware.InitMiddleware(), c)
	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(list); err != nil {
		log.Fatal(err)
	}
}
