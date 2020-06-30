package main

import (
	"context"
	"net"
	"time"

	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/ezio1119/fishapp-chat/infrastructure"
	"github.com/ezio1119/fishapp-chat/infrastructure/middleware"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers"
	"github.com/ezio1119/fishapp-chat/interfaces/repo"
	"github.com/ezio1119/fishapp-chat/pb"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	dbConn := infrastructure.NewGormConn()
	defer dbConn.Close()

	rClient := infrastructure.NewRedisClient()
	defer rClient.Close()

	natsConn := infrastructure.NewNatsStreamingConn()
	defer natsConn.Close()

	grpcConn, err := grpc.DialContext(ctx, conf.C.API.ImageURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	imageC := pb.NewImageServiceClient(grpcConn)
	imageRepo := repo.NewImageRepository(imageC)

	timeOut := time.Duration(conf.C.Sv.Timeout) * time.Second

	chatC := controllers.NewChatController(
		interactor.NewChatInteractor(
			dbConn,
			rClient,
			imageRepo,
			timeOut,
		),
	)

	eventC := controllers.NewEventController(interactor.NewEventInteractor(dbConn, imageRepo, timeOut))
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
