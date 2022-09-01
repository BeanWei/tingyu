package g

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

var (
	pool     *ants.Pool
	poolOnce sync.Once
)

func Pool() *ants.Pool {
	poolOnce.Do(func() {
		pool, _ = ants.NewPool(ants.DefaultAntsPoolSize)
	})
	return pool
}
