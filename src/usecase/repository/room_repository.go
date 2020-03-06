package repository

import (
	"context"

	"github.com/ezio1119/fishapp-chat/domain"
)

type RoomRepository interface {
	CreateRoomAddMember(ctx context.Context, rm *domain.Room, m *domain.Member) error
	GetRoomByID(ctx context.Context, id int64) (*domain.Room, error)
	GetMemberByUserIDAndRoomID(ctx context.Context, uID int64, rID int64) (*domain.Member, error)
	AddMember(ctx context.Context, m *domain.Member) error
	CreateMessage(ctx context.Context, m *domain.Message) error
}
