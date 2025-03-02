package student

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IStudentWorker interface {
	Start(kurl string)
	HandleUserCreatedEvent(ctx context.Context, msg kafka.Message) error
}
