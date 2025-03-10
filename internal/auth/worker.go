package auth

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IAuthWorker interface {
	Start(kurl string)
	HandleCreateUserEvent(ctx context.Context, msg kafka.Message) error
}
