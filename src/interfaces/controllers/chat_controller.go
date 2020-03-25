package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/sync/errgroup"
)

type chatController struct {
	chatInteractor interactor.ChatInteractor
}

func NewChatController(ci interactor.ChatInteractor) chat_grpc.ChatServiceServer {
	return &chatController{ci}
}

func (c *chatController) CreateRoom(ctx context.Context, in *chat_grpc.CreateRoomReq) (*chat_grpc.Room, error) {
	r := &domain.Room{
		PostID: in.PostId,
		Members: []*domain.Member{
			{UserID: in.UserId},
		},
	}
	if err := c.chatInteractor.CreateRoom(ctx, r); err != nil {
		return nil, err
	}
	return convRoomProto(r)
}

func (c *chatController) GetRoom(ctx context.Context, in *chat_grpc.GetRoomReq) (*chat_grpc.Room, error) {
	r, err := c.chatInteractor.GetRoom(ctx, in.PostId, in.UserId)
	if err != nil {
		return nil, err
	}
	return convRoomProto(r)
}

func (c *chatController) ListMembers(ctx context.Context, in *chat_grpc.ListMembersReq) (*chat_grpc.ListMembersRes, error) {
	list, err := c.chatInteractor.ListMembers(ctx, in.RoomId, in.UserId)
	if err != nil {
		return nil, err
	}
	listMProto, err := convListMembersProto(list)
	if err != nil {
		return nil, err
	}
	return &chat_grpc.ListMembersRes{Members: listMProto}, nil
}

func (c *chatController) CreateMember(ctx context.Context, in *chat_grpc.CreateMemberReq) (*chat_grpc.Member, error) {
	m := &domain.Member{RoomID: in.RoomId, UserID: in.UserId}
	if err := c.chatInteractor.CreateMember(ctx, m); err != nil {
		return nil, err
	}
	return convMemberProto(m)
}

func (c *chatController) DeleteMember(ctx context.Context, in *chat_grpc.DeleteMemberReq) (*empty.Empty, error) {
	if err := c.chatInteractor.DeleteMember(ctx, in.RoomId, in.UserId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *chatController) ListMessages(ctx context.Context, in *chat_grpc.ListMessagesReq) (*chat_grpc.ListMessagesRes, error) {
	list, err := c.chatInteractor.ListMessages(ctx, in.RoomId, in.UserId)
	if err != nil {
		return nil, err
	}
	listMProto, err := convListMessagesProto(list)
	if err != nil {
		return nil, err
	}
	return &chat_grpc.ListMessagesRes{Messages: listMProto}, nil
}

func (c *chatController) CreateMessage(ctx context.Context, in *chat_grpc.CreateMessageReq) (*chat_grpc.Message, error) {
	m := &domain.Message{
		Body:   in.Body,
		RoomID: in.RoomId,
		UserID: in.UserId,
	}
	if err := c.chatInteractor.CreateMessage(ctx, m); err != nil {
		return nil, err
	}
	return convMessageProto(m)
}

func (c *chatController) StreamMessage(in *chat_grpc.StreamMessageReq, stream chat_grpc.ChatService_StreamMessageServer) error {
	eg, ctx := errgroup.WithContext(stream.Context())
	msgChan := make(chan *domain.Message)
	go func() {
		eg.Wait()
		close(msgChan)
	}()

	eg.Go(func() error {
		if err := c.chatInteractor.StreamMessage(ctx, in.RoomId, in.UserId, msgChan); err != nil {
			return err
		}
		return nil
	})
	for m := range msgChan {
		mProto, err := convMessageProto(m)
		if err != nil {
			return err
		}
		if err := stream.Send(mProto); err != nil {
			return err
		}
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
