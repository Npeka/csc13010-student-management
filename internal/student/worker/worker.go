package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/student"
	kafkas "github.com/csc13010-student-management/pkg/kafka"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type studentWorker struct {
	su student.IStudentUsecase
	lg *logger.LoggerZap
}

func NewStudentWorker(
	su student.IStudentUsecase,
	lg *logger.LoggerZap,
) student.IStudentWorker {
	return &studentWorker{
		su: su,
		lg: lg,
	}
}

func (sw *studentWorker) Start(kurl string) {
	krs := []kafkas.KafkaReader{
		{
			Topic:   events.AuthUserCreated,
			GroupID: fmt.Sprintf("%v.g", events.AuthCreateUser),
			Handler: sw.HandleUserCreatedEvent,
			MaxIns:  1,
		},
	}
	kafkas.StartKafkaConsumers(kurl, krs)
}

func (sw *studentWorker) HandleUserCreatedEvent(ctx context.Context, msg kafka.Message) error {
	var event events.UserCreatedEvent
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	if err := sw.su.UpdateUserIDByUsername(ctx, event.Username, event.UserID); err != nil {
		return err
	}

	return nil
}
