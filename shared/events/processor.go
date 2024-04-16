package events

import (
	"context"
	db "shared/db/sqlc"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

type TaskProcessor interface {
	Start(isWorker bool) error
	ProcessPortScanner(ctx context.Context, task *asynq.Task) error
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
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("type", task.Type()).
					Bytes("payload", task.Payload()).Msg("process task failed")
			}),
			Logger: NewLogger(),
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start(isWorker bool) error {
	InitTaskDistributor()
	asynqMux := asynq.NewServeMux()
	if isWorker {
		asynqMux.HandleFunc(string(PortScanner), processor.ProcessPortScanner)
	} else {
		asynqMux.HandleFunc(string(ScannerResult), processor.ProcessPortScannerResponse)
	}

	return processor.server.Start(asynqMux)
}
