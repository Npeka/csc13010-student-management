package auth

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IAuthWorker interface {
	Start(kurl string)
	HandleStudentCreatedEvent(ctx context.Context, msg kafka.Message) error
}
