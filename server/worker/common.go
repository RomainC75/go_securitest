package worker

import (
	"fmt"
	"server/config"

	"github.com/hibiken/asynq"
)

var taskDistributor TaskDistributor

func Init() {
	config := config.Get()
	fmt.Printf("task distributo init : %s:%s\n", config.Redis.Host, config.Redis.Port)
	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
	}
	taskDistributor = NewRedisTaskDistributor(redisOpt)
}

func Get() TaskDistributor {
	return taskDistributor
}
