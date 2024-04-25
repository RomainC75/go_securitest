package events

import (
	"context"
	"encoding/json"
	"fmt"
	"server/event_logic"
	"server/utils"
	"shared/scenarios"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (processor *RedisTaskProcessor) ProcessPortScannerResponse(ctx context.Context, task *asynq.Task) error {
	fmt.Println("===========GET RESPONSE================")
	var payload scenarios.ScanResult
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload : %w", asynq.SkipRetry)
	}
	utils.PrettyDisplay("worker response", payload)

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("targetIp", "lkj").Msg("PROCESSED task")

	event_logic.HandleScanResponse(payload)

	return nil
}
