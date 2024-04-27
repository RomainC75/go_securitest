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
	var originalPayload work_dto.PortTestScenario
	if err := json.Unmarshal(task.Payload(), &originalPayload); err != nil {
		return fmt.Errorf("failed to unmarshal originalPayload : %w", asynq.SkipRetry)
	}
	targetIp := originalPayload.IPRange.IpMin
	// user, err := processor.store.GetUserByEmail(ctx, originalPayload.Username)

	// TODO: send email to user

	// result

	result, err := scenarios.Scan(targetIp, originalPayload.PortRange.Min, originalPayload.PortRange.Max)
	fmt.Printf("===========FINISHED ================")
	if err != nil {
		log.Error().Str("scenario Error : ", err.Error())
	}
	// strResult, err := scenarios.GetString(result)

	log.Info().Str("type", task.Type()).Bytes("originalPayload", task.Payload()).
		Str("targetIp", targetIp).Msg("PROCESSED task")
	//============================================
	distributor := Get()
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		// asynq.ProcessIn(time.Second),
		// + send to other queue :-)
		asynq.Queue(string(CriticalQueueRes)),
	}
	fmt.Println("====> distributor : ", distributor)
	err = distributor.DistributeTaskSendWorkBack(
		ctx,
		&result,
		&originalPayload,
		PortScanner,
		opts...,
	)

	return nil
}
