package main

import (
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
	defer dbConn.Close()

	rClient := infrastructure.NewRedisClient()
	defer rClient.Close()

	natsConn := infrastructure.NewNatsStreamingConn()
	defer natsConn.Close()

	timeOut := time.Duration(conf.C.Sv.Timeout) * time.Second

	chatC := controllers.NewChatController(
		interactor.NewChatInteractor(
			dbConn,
			rClient,
			timeOut,
		),
	)

	eventC := controllers.NewEventController(interactor.NewEventInteractor(dbConn, timeOut))
	server := infrastructure.NewGrpcServer(middleware.InitMiddleware(), chatC)

	if err := infrastructure.StartSubscribeNats(eventC, natsConn); err != nil {
		panic(err)
	}

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		panic(err)
	}

	if err = server.Serve(list); err != nil {
		panic(err)
	}
}
