package g

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdb     *redis.Client
	rdbOnce sync.Once
)

func Redis() *redis.Client {
	rdbOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     Cfg().Redis.Addr,
			Password: Cfg().Redis.Pass,
			DB:       Cfg().Redis.DB,
		})
	})
	return rdb
}
