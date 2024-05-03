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

func (processor *RedisTaskProcessor) ProcessPortScanner(ctx context.Context, task *asynq.Task) error {
	var originalPayload work_dto.PortTestScenarioRequest
	if err := json.Unmarshal(task.Payload(), &originalPayload); err != nil {
		return fmt.Errorf("failed to unmarshal originalPayload : %w", asynq.SkipRetry)
	}

	result, err := scenarios.Scan(originalPayload)
	if err != nil {
		log.Error().Str("scenario Error : ", err.Error())
	}

	log.Info().Str("type", task.Type()).Bytes("originalPayload", task.Payload()).
		Str("targetIp", originalPayload.IPRange.IpMin).Str("targetIp", originalPayload.IPRange.IpMax).Msg("PROCESSED task")
	distributor := Get()
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		// asynq.ProcessIn(time.Second),
		// + send to other queue :-)
		asynq.Queue(string(CriticalQueueRes)),
	}
	err = distributor.DistributeTaskSendWorkBack(
		ctx,
		&result,
		PortScanner,
		opts...,
	)
	return err
}
