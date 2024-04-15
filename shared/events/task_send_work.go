package events

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	request_dto "server/api/dto/request"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type ScenarioSelector string

const (
	PortScanner ScenarioSelector = "task:port_scanner"
)

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendWork(
	ctx context.Context,
	payload *request_dto.PortTestScenario,
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

func (processor *RedisTaskProcessor) ProcessPortScanner(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload : %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByEmail(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exist: %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	// TODO: send email to user
	fmt.Println("===> NEW TASK : ")
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("PROCESSED task")
	return nil
}
