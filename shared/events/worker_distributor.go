package events

import (
	"context"
	"encoding/json"
	"fmt"
	work_dto "shared/dto"
	"shared/scenarios"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (distributor *RedisTaskDistributor) DistributeTaskSendWorkBack(
	ctx context.Context,
	payload *scenarios.ScanResult,
	originalPayload *work_dto.PortTestScenario,
	scenario ScenarioSelector,
	opts ...asynq.Option,
) error {
	// TODO : merge payload & originalPayload
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("fails to marshal payload: %w", err)
	}

	task := asynq.NewTask(string(ScannerResult), jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("fails to enqueue task : %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("ENQUEUED task")

	return nil
}
