package infrastructure

import (
	"github.com/ezio1119/fishapp-chat/infrastructure/middleware"
	"github.com/ezio1119/fishapp-chat/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(middLe *middleware.Middleware, chatController pb.ChatServiceServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middLe.UnaryLogingInterceptor(),
			middLe.UnaryValidationInterceptor(),
			middLe.UnaryRecoveryInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			middLe.StreamLogingInterceptor(),
			middLe.StreamValidationInterceptor(),
			middLe.StreamRecoveryInterceptor(),
		),
	)

	pb.RegisterChatServiceServer(server, chatController)
	reflection.Register(server)
	return server
}
