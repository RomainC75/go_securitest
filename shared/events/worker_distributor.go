package events

import (
	"context"
	"encoding/json"
	"fmt"
	"shared/scenarios"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (distributor *RedisTaskDistributor) DistributeTaskSendWorkBack(
	ctx context.Context,
	payload *scenarios.ScanResult,
	scenario ScenarioSelector,
	opts ...asynq.Option,
) error {
	fmt.Println("=====>DISTRIBUTE BACK")
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("faild to marshal payload: %w", err)
	}

	fmt.Println("222222222222222222222222")
	task := asynq.NewTask(string(ScannerResult), jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("fails to enqueue task : %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("ENQUEUED task")

	return nil
}
