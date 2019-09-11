package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

// 使用文档 : {@link https://github.com/go-redis/redis}

func TestRedisPool(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr : "localhost:6379",
	})

	err := client.Set("key", 10086, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}