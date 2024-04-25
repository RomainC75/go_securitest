package events

import (
	"context"
	work_dto "shared/dto"
	"shared/scenarios"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendWork(
		ctx context.Context,
		payload *work_dto.PortTestScenario,
		scenario ScenarioSelector,
		opts ...asynq.Option,
	) error
	DistributeTaskSendWorkBack(
		ctx context.Context,
		payload *scenarios.ScanResult,
		originalPayload *work_dto.PortTestScenario,
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
