package events

import (
	"context"
	"encoding/json"
	"fmt"
	work_dto "shared/dto"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (distributor *RedisTaskDistributor) DistributeTaskSendWork(
	ctx context.Context,
	payload *work_dto.PortTestScenarioRequest,
	scenario ScenarioSelector,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("faild to marshal payload: %w", err)
	}

	task := asynq.NewTask(string(PortScanner), jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("fails to enqueue task : %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("ENQUEUED task")

	return nil
}
