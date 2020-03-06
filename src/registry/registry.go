package registry

import (
	"time"

	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db         *gorm.DB
	kvs        *redis.Client
	ctxTimeout time.Duration
}

type Registry interface {
	NewChatController() chat_grpc.ChatServiceServer
}

func NewRegistry(db *gorm.DB, kvs *redis.Client, t time.Duration) Registry {
	return &registry{db, kvs, t}
}
