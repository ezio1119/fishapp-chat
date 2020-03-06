package repository

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/usecase/repository"
	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queuingRepository struct {
	sync.RWMutex
	client      *redis.Client
	blockingIDs map[string]int64
}

func NewQueuingRepository(r *redis.Client) repository.QueuingRepository {
	return &queuingRepository{sync.RWMutex{}, r, map[string]int64{}}
}

func (r *queuingRepository) XADD(m *domain.Message) error {
	rID := strconv.FormatInt(m.RoomID, 10)
	res := r.client.XAdd(&redis.XAddArgs{
		Stream: rID,
		Values: map[string]interface{}{
			"message": m,
		},
		MaxLen: 100,
	})
	if err := res.Err(); err != nil {
		return err
	}
	return nil
}

func (r *queuingRepository) GetBlockingID(uID string) int64 {
	r.RLock()
	defer r.RUnlock()
	return r.blockingIDs[uID]
}

func (r *queuingRepository) DeleteBlockingID(uID string) {
	r.Lock()
	defer r.Unlock()
	delete(r.blockingIDs, uID)
	log.Println("blockingIDs: ", r.blockingIDs)
}

func (r *queuingRepository) XREADGROUP(ctx context.Context, g string, s []string) ([]*domain.Message, error) {
	r.Lock()
	r.blockingIDs[g] = r.client.ClientID().Val()
	r.Unlock()

	streams, err := r.client.WithContext(ctx).XReadGroup(&redis.XReadGroupArgs{
		Group:   g,
		Streams: s,
		Block:   0,
	}).Result()
	if err != nil {
		return nil, err
	}

	msgs := []*domain.Message{}
	for _, s := range streams {
		for _, xmsg := range s.Messages {
			log.Printf("%#v", xmsg)
			v, ok := xmsg.Values["message"]
			if !ok {
				return nil, status.Errorf(codes.Internal, "redis streams: message key not found")
			}
			msg, ok := v.(string)
			if !ok {
				return nil, status.Errorf(codes.Internal, "redis streams: type assertion on string failed")
			}
			m := &domain.Message{}
			if err := m.UnmarshalBinary([]byte(msg)); err != nil {
				return nil, err
			}
			msgs = append(msgs, m)
		}
	}
	return msgs, nil
}

func (r *queuingRepository) XGroupCreateMkStream(s, g string) error {
	if err := r.client.XGroupCreateMkStream(s, g, "$").Err(); err != nil {
		return err
	}
	return nil
}

func (r *queuingRepository) ClientUnblock(id int64) error {
	return r.client.ClientUnblock(id).Err()
}
