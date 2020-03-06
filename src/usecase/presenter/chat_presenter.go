package presenter

import (
	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
)

type ChatPresenter interface {
	TransformRoomProto(rm *domain.Room) (*chat_grpc.Room, error)
	TransformMemberProto(m *domain.Member) (*chat_grpc.Member, error)
	TransformMessageProto(m *domain.Message) (*chat_grpc.Message, error)
	StreamMessage(m *domain.Message, stream chat_grpc.ChatService_StreamMessageServer) error
}
