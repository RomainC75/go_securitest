package events

import (
	"context"
	request_dto "server/api/dto/request"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendWork(
		ctx context.Context,
		payload *request_dto.PortTestScenario,
		scenario ScenarioSelector,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}

// 9:36
