package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/ezio1119/fishapp-chat/infrastructure"
	"github.com/ezio1119/fishapp-chat/infrastructure/middleware"
	"github.com/ezio1119/fishapp-chat/registry"
)

func main() {
	dbConn := infrastructure.NewGormConn()
	rConn := infrastructure.NewRedisClient()
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	t := time.Duration(conf.C.Sv.Timeout) * time.Second
	registry := registry.NewRegistry(dbConn, rConn, t)
	chatController := registry.NewChatController()

	middLe := middleware.InitMiddleware()
	server := infrastructure.NewGrpcServer(middLe, chatController)
	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	if err = server.Serve(list); err != nil {
		log.Fatal(err)
	}
}
