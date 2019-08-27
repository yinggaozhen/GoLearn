package redis

import (
	"fmt"
	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
	"testing"
	"time"
)

type Object struct {
	Str string
	Num int
}

// 使用文档 : {@link https://github.com/go-redis/cache}

// FIXME 这里有一个能成为 go-redis/redis contributors 的机会

func TestRingRedis(t *testing.T) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": ":6379",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	codec.UseLocalCache(1, 100 * time.Second)

	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	err := codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour,
	})

	if err != nil {
		panic(err)
	}

	var wanted Object
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println(wanted)
	}
	codec.Get(key, &wanted)
	codec.Get(key, &wanted)
	codec.Get(key, &wanted)

	// Output: {mystring 42}
}