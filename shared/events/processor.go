package events

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type WorkRequestTaskName string

const (
	CriticalQueueReq WorkRequestTaskName = "critical_req"
	DefaultQueueReq  WorkRequestTaskName = "default_req"
	LowQueueReq      WorkRequestTaskName = "low_req"
)

type WorkResponseTaskName string

const (
	CriticalQueueRes WorkResponseTaskName = "critical_res"
	DefaultQueueRes  WorkResponseTaskName = "default_res"
	LowQueueRes      WorkResponseTaskName = "low_res"
)

type TaskProcessor interface {
	Start(isWorker bool) error
	ProcessPortScanner(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
}

func getQueues(isWorker bool) map[string]int {

	if isWorker {
		return map[string]int{
			string(CriticalQueueReq): 6,
			string(DefaultQueueReq):  3,
			string(LowQueueReq):      1,
		}
	} else {
		return map[string]int{
			string(CriticalQueueRes): 6,
			string(DefaultQueueRes):  3,
			string(LowQueueRes):      1,
		}
	}
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, isWorker bool) TaskProcessor {

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			// accepted queue names !
			Queues: getQueues(isWorker),
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("type", task.Type()).
					Bytes("payload", task.Payload()).Msg("process task failed")
			}),
			Logger: NewLogger(),
		},
	)

	return &RedisTaskProcessor{
		server: server,
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
