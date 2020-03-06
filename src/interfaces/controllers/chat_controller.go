package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
)

type chatController struct {
	chatInteractor interactor.ChatInteractor
}

func NewChatController(ci interactor.ChatInteractor) chat_grpc.ChatServiceServer {
	return &chatController{ci}
}

func (c *chatController) CreateChatRoom(ctx context.Context, in *chat_grpc.CreateRoomReq) (*chat_grpc.Room, error) {
	r, err := c.chatInteractor.CreateChatRoom(ctx, &domain.Room{
		PostID: in.PostId,
	}, &domain.Member{
		UserID: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *chatController) AddMember(ctx context.Context, in *chat_grpc.AddMemberReq) (*chat_grpc.Member, error) {
	m, err := c.chatInteractor.AddMember(ctx, &domain.Member{
		RoomID: in.RoomId,
		UserID: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatController) SendMessage(ctx context.Context, in *chat_grpc.SendMessageReq) (*chat_grpc.Message, error) {
	m, err := c.chatInteractor.SendMessage(ctx, &domain.Message{
		Body:   in.Body,
		RoomID: in.RoomId,
		UserID: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatController) StreamMessage(in *chat_grpc.StreamMessageReq, s chat_grpc.ChatService_StreamMessageServer) error {
	if err := c.chatInteractor.StreamMessage(in.RoomIds, in.UserId, s); err != nil {
		return err
	}
	return nil
}
