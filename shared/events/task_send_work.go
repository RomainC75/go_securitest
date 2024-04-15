package events

import (
	"context"
	"encoding/json"
	"fmt"
	request_dto "server/api/dto/request"
	"shared/scenarios"

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
	var payload request_dto.PortTestScenario
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload : %w", asynq.SkipRetry)
	}
	targetIp := payload.IPRange.IpMin
	// user, err := processor.store.GetUserByEmail(ctx, payload.Username)

	// TODO: send email to user

	result, err := scenarios.Scan(targetIp, payload.PortRange.Min, payload.PortRange.Max)
	fmt.Printf("===========FINISHED ================")
	if err != nil {
		log.Error().Str("scenario Error : ", err.Error())
	}
	log.Warn().Str("=====> scenario Ress : ", scenarios.GetString(result))
	fmt.Println("===> DONE TASKS : ", scenarios.GetString(result))
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("targetIp", targetIp).Msg("PROCESSED task")
	return nil
}
