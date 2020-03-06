package presenter

import (
	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/ezio1119/fishapp-chat/usecase/presenter"
	"github.com/golang/protobuf/ptypes"
)

type chatPresenter struct {
}

func NewChatPresenter() presenter.ChatPresenter {
	return &chatPresenter{}
}

func (p *chatPresenter) TransformRoomProto(rm *domain.Room) (*chat_grpc.Room, error) {
	updatedAt, err := ptypes.TimestampProto(rm.UpdatedAt)
	if err != nil {
		return nil, domain.WrapOnChatPresenErr(err)
	}
	createdAt, err := ptypes.TimestampProto(rm.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &chat_grpc.Room{
		Id:        rm.ID,
		PostId:    rm.PostID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (p *chatPresenter) TransformMemberProto(m *domain.Member) (*chat_grpc.Member, error) {
	updatedAt, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		return nil, domain.WrapOnChatPresenErr(err)
	}
	createdAt, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		return nil, domain.WrapOnChatPresenErr(err)
	}
	return &chat_grpc.Member{
		Id:        m.ID,
		RoomId:    m.RoomID,
		UserId:    m.UserID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (p *chatPresenter) TransformMessageProto(m *domain.Message) (*chat_grpc.Message, error) {
	updatedAt, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		return nil, domain.WrapOnChatPresenErr(err)
	}
	createdAt, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		return nil, domain.WrapOnChatPresenErr(err)
	}
	return &chat_grpc.Message{
		Id:        m.ID,
		Body:      m.Body,
		RoomId:    m.RoomID,
		UserId:    m.UserID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (p *chatPresenter) StreamMessage(m *domain.Message, s chat_grpc.ChatService_StreamMessageServer) error {
	mP, err := p.TransformMessageProto(m)
	if err != nil {
		return err
	}
	if err := s.Send(mP); err != nil {
		return err
	}
	return nil
}
