package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

// 使用文档 : {@link https://github.com/go-redis/redis}

func TestClusterRedis(t *testing.T) {
	addrs := []string{
		"172.17.0.2:7000",
		"172.17.0.3:7000",
		"172.17.0.4:7000",
		"172.17.0.5:7000",
		"172.17.0.6:7000",
		"172.17.0.7:7000",
	}

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs : addrs,
	})
	err := client.Set("key", "value", 0).Err()
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