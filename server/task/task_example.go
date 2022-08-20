package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const TypeExample = "example"

type examplePayload struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func NewExampleTask(payload *examplePayload) *asynq.Task {
	tp, _ := json.Marshal(payload)
	return asynq.NewTask(TypeExample, tp)
}

func HandleExampleTask(ctx context.Context, t *asynq.Task) error {
	var payload *examplePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf(`
		# Test Asynq
		Subject: %s
		Message: %s
		Time: %s
	`, payload.Subject, payload.Message, time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
