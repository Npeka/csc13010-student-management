package auditlog

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IAuditLogWorker interface {
	Start(kurl string)
	HandleTableChangedEvent(ctx context.Context, msg kafka.Message) error
}
