package infrastructure

import (
	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

type EventController interface {
	CreateRoom(m *stan.Msg)
}

func NewNatsStreamingConn() stan.Conn {
	conn, err := stan.Connect(conf.C.Nats.ClusterID, "fishapp-chat-"+uuid.New().String(), stan.NatsURL(conf.C.Nats.URL))
	if err != nil {
		panic(err)
	}
	return conn
}

func StartSubscribeNats(c EventController, conn stan.Conn) error {
	_, err := conn.QueueSubscribe("create.room", conf.C.Nats.QueueGroup, c.CreateRoom, stan.DurableName(conf.C.Nats.QueueGroup))

	if err != nil {
		return err
	}

	return nil
}

// func StartSubscribeNats(c EventController, conn stan.Conn) error {
// 	_, err := conn.QueueSubscribe("create.room", conf.C.Nats.QueueGroup, func(m *stan.Msg) {
// 		e := &event.Event{}
// 		if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
// 			log.Printf("error wrong subject data type : %s", err)
// 		}

// 		switch e.EventType {
// 		case "create.room":
// 			data := &event.CreateRoom{}
// 			if err := protojson.Unmarshal(e.EventData, e); err != nil {
// 				log.Printf("error wrong eventdata type: %s", err)
// 			}
// 			if err := c.CreateRoom(context.Background(), data); err != nil {
// 				log.Println(err)
// 			}
// 		}

// 	}, stan.DurableName(conf.C.Nats.QueueGroup))

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
