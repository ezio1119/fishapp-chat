package repository

import (
	"context"

	"github.com/ezio1119/fishapp-chat/domain"
)

type QueuingRepository interface {
	GetBlockingID(userID string) int64
	DeleteBlockingID(userID string)
	ClientUnblock(clientID int64) error
	XADD(m *domain.Message) error
	XREADGROUP(ctx context.Context, group string, streams []string) ([]*domain.Message, error)
	XGroupCreateMkStream(stream, group string) error
}
