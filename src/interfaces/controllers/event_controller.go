package controllers

import (
	"context"
	"log"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb/event"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/encoding/protojson"
)

type eventController struct {
	eventInteractor interactor.EventInteractor
}

func NewEventController(ei interactor.EventInteractor) *eventController {
	return &eventController{ei}
}

func (c *eventController) CreateRoom(m *stan.Msg) {
	ctx := context.Background()

	e := &event.Event{}
	if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
		log.Printf("error wrong subject data type : %s", err)
		return
	}

	eventData := &event.CreateRoom{}
	if err := protojson.Unmarshal(e.EventData, eventData); err != nil {
		log.Printf("error wrong eventdata type: %s", err)
		return
	}

	if err := c.eventInteractor.CreateRoom(ctx, &domain.Room{
		PostID: eventData.PostId,
		Members: []*domain.Member{
			{UserID: eventData.UserId},
		},
	}, eventData.SagaId); err != nil {
		log.Println(err)
		return
	}

}
