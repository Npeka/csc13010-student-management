package notification

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type INotificationWorker interface {
	Start(kurl string)
	HandleNotiCreateEvent(ctx context.Context, msg kafka.Message) error
}
