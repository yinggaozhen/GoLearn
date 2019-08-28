package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

// 使用文档 : {@link https://github.com/go-redis/redis}

func TestRedis(t *testing.T) {
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

// FIXME 这里有一个能成为 go-redis/redis contributors 的机会
func TestRedisDoCommand(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr : "localhost:6379",
	})

	key := "redis_test_key"

	client.Do("set", key, 1)

	return1 := client.Do("ttl", key).Val().(int64)
	fmt.Println(return1)

	if return1 == -1 {
		return2 := client.Do("expire", key, int64(60)).Val()
		fmt.Println(return2)

		return3 := client.Do("ttl", key).Val().(int64)
		fmt.Println(return3)
	}
}