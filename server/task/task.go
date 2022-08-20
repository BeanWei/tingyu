package task

import (
	"github.com/BeanWei/tingyu/g"
	"github.com/hibiken/asynq"
)

var client *asynq.Client

func Run() {
	cfg := g.Cfg()
	rdb := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Pass,
		DB:       cfg.Redis.DB,
	}
	client = asynq.NewClient(rdb)
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeExample, HandleExampleTask)
	if err := asynq.NewServer(rdb, asynq.Config{}).Run(mux); err != nil {
		panic(err)
	}
}

func Client() *asynq.Client {
	return client
}
