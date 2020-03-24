package controllers

import (
	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/golang/protobuf/ptypes"
)

func convRoomProto(r *domain.Room) (*chat_grpc.Room, error) {
	uAt, err := ptypes.TimestampProto(r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	cAt, err := ptypes.TimestampProto(r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &chat_grpc.Room{
		Id:        r.ID,
		PostId:    r.PostID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convMemberProto(m *domain.Member) (*chat_grpc.Member, error) {
	uAt, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	cAt, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &chat_grpc.Member{
		Id:        m.ID,
		RoomId:    m.RoomID,
		UserId:    m.UserID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convListMembersProto(list []*domain.Member) ([]*chat_grpc.Member, error) {
	listM := make([]*chat_grpc.Member, len(list))
	for i, m := range list {
		mProto, err := convMemberProto(m)
		if err != nil {
			return nil, err
		}
		listM[i] = mProto
	}
	return listM, nil
}

func convMessageProto(m *domain.Message) (*chat_grpc.Message, error) {
	uAt, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	cAt, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &chat_grpc.Message{
		Id:        m.ID,
		Body:      m.Body,
		RoomId:    m.RoomID,
		UserId:    m.UserID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convListMessagesProto(list []*domain.Message) ([]*chat_grpc.Message, error) {
	listM := make([]*chat_grpc.Message, len(list))
	for i, m := range list {
		mProto, err := convMessageProto(m)
		if err != nil {
			return nil, err
		}
		listM[i] = mProto
	}
	return listM, nil
}
