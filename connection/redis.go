package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-10914.c98.us-east-1-4.ec2.redns.redis-cloud.com:10914",
		Username: "default",
		Password: "5qrVRqQytOg0OQLvDaRIkSmhUQpMqAJH",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Ping to redis server failed: %v", err)
	}

	fmt.Println("Redis connection successfully")

	return rdb
}
