package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb/event"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/encoding/protojson"
)

type eventInteractor struct {
	db         *gorm.DB
	ctxTimeout time.Duration
}

func NewEventInteractor(db *gorm.DB, t time.Duration) EventInteractor {
	return &eventInteractor{db, t}
}

type EventInteractor interface {
	CreateRoom(ctx context.Context, r *domain.Room, sagaID string) error
}

func (i *eventInteractor) CreateRoom(ctx context.Context, r *domain.Room, sagaID string) error {
	err := i.db.Create(r).Error
	if err != nil {

		eventDataJson, err := protojson.Marshal(&event.CreateRoomFailed{
			SagaId:  sagaID,
			Message: err.Error(),
		})
		if err != nil {
			return err
		}
		event := &domain.Outbox{EventType: "create.room.failed", EventData: eventDataJson}

		if err := i.db.Create(event).Error; err != nil {
			return err
		}
	}

	rProto, err := convRoomProto(r)
	if err != nil {
		return err
	}

	eventDataJson, err := protojson.Marshal(&event.RoomCreated{
		SagaId: sagaID,
		Room:   rProto,
	})
	if err != nil {
		return err
	}
	event := &domain.Outbox{EventType: "room.created", EventData: eventDataJson}

	if err := i.db.Create(event).Error; err != nil {
		return err
	}
	return nil
}
