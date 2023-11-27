package initial

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-blog/global"
)

func Redis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
		PoolSize: redisCfg.PoolSize,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Errorf("redis connect ping failed, err: %v", err)
		return nil
	} else {
		global.Logger.Infof("redis connect ping response: %v", pong)
		return client
	}
}
