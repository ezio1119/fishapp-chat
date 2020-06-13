package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

type EventController interface {
	CreateRoom(m *stan.Msg)
	PostDeleted(m *stan.Msg)
}

func NewNatsStreamingConn() stan.Conn {
	clientID := uuid.New().String()
	log.Printf("nats clientID is %s", clientID)
	conn, err := stan.Connect(conf.C.Nats.ClusterID, "fishapp-chat-"+uuid.New().String(), stan.NatsURL(conf.C.Nats.URL))
	if err != nil {
		panic(err)
	}
	return conn
}

func StartSubscribeNats(c EventController, conn stan.Conn) error {
	_, err := conn.QueueSubscribe("create.room", conf.C.Nats.QueueGroup, c.CreateRoom, stan.DurableName(conf.C.Nats.QueueGroup))
	_, err = conn.QueueSubscribe("post.deleted", conf.C.Nats.QueueGroup, c.PostDeleted, stan.DurableName(conf.C.Nats.QueueGroup), stan.SetManualAckMode())

	if err != nil {
		return err
	}

	return nil
}

// func StartSubscribeNats(c EventController, conn stan.Conn) error {
// 	_, err := conn.QueueSubscribe("create.room", conf.C.Nats.QueueGroup, func(m *stan.Msg) {
// 		e := &pb.Event{}
// 		if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
// 			log.Printf("error wrong subject data type : %s", err)
// 		}

// 		switch e.EventType {
// 		case "create.room":
// 			data := &pb.CreateRoom{}
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
