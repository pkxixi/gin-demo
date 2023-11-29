package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestBlockUtil(t *testing.T) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "test123456",
		DB:       0,
		PoolSize: 10,
	})
	ctx := context.Background()
	// Get non-exist Key
	v, _ := redisCli.Get(ctx, "LIMIT_XXXX").Result()
	fmt.Println("v", v)

	v2, err := redisCli.Set(ctx, "LIMIT_KE_PAN", 1, time.Minute).Result()
	fmt.Println("v2", v2, err)
	r, err := redisCli.Get(ctx, "LIMIT_KE_PAN").Result()
	fmt.Printf("r type %T, r value %v, err content %v\n", r, r, err)
	v3, _ := redisCli.Incr(ctx, "LIMIT_KE_PAN").Result()
	fmt.Println("v3", v3) // v3 2
}
