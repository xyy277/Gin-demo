package initialize

import (
	"context"

	"gsafety/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GS_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GS_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GS_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GS_REDIS = client
	}
}
