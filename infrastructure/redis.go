package infrastructure

import (
	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.C.Kvs.Host + ":" + conf.C.Kvs.Port,
		Password: conf.C.Kvs.Pass,
		DB:       conf.C.Kvs.Db,
		Network:  conf.C.Kvs.Net,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
