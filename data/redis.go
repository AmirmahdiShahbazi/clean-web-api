package data

import (
	"context"
	"time"

	"github.com/AmirmahdiShahbazi/clean-web-api/config"
	"github.com/AmirmahdiShahbazi/clean-web-api/pkg/logger"
	"github.com/go-redis/redis/v8"
)
var RedisClient *redis.Client
func ConnectToRedis(){
	redisConfigs := config.GetRedisConfigs()
	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisConfigs.Addr,
		Password: redisConfigs.Password,
		DB: redisConfigs.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)  
	defer cancel()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {  
		logger.NewLogger().Fatal(logger.Redis, logger.Startup, err.Error(), nil)
	}  
}

func GetRedisClient() *redis.Client{
	return RedisClient
}

