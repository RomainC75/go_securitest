package worker

import (
	"context"
	db "server/db/sqlc"

	"github.com/hibiken/asynq"
)

const (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			// accepted queue names !
			Queues: map[string]int{
				CriticalQueue: 6,
				DefaultQueue:  3,
				LowQueue:      1,
			},
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	asynqMux := asynq.NewServeMux()
	asynqMux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(asynqMux)
}
