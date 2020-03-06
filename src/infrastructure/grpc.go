package infrastructure

import (
	"github.com/ezio1119/fishapp-chat/infrastructure/middleware"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(middLe middleware.Middleware, chatController chat_grpc.ChatServiceServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middLe.LoggerInterceptor(),
			middLe.ValidatorInterceptor(),
			middLe.RecoveryInterceptor(),
		)),
	)
	chat_grpc.RegisterChatServiceServer(server, chatController)
	reflection.Register(server)
	return server
}
