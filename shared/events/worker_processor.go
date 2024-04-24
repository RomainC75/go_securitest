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
	var payload work_dto.PortTestScenario
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
	// strResult, err := scenarios.GetString(result)

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
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
		PortScanner,
		opts...,
	)

	return nil
}
