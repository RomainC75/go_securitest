package events

import (
	"fmt"
	"shared/config"
	db "shared/db/sqlc"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

var taskDistributor TaskDistributor

func InitTaskDistributor() {
	config := config.Get()
	fmt.Printf("task distributo init : %s:%s\n", config.Redis.Host, config.Redis.Port)
	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
	}

	taskDistributor = NewRedisTaskDistributor(redisOpt)
}

func InitTaskProcessor(store db.Store) {
	// supposed to run as a go routine
	config := config.Get()
	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
	}
	taskProcessor := NewRedisTaskProcessor(redisOpt, store)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func Get() TaskDistributor {
	return taskDistributor
}
