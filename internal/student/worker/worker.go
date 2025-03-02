package worker

import (
	"context"
	"encoding/json"

	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/initialize"
	"github.com/csc13010-student-management/internal/student"
	kafkaUtils "github.com/csc13010-student-management/pkg/kafka"
	"github.com/segmentio/kafka-go"
)

type studentWorker struct {
	su student.IStudentUsecase
}

func NewStudentWorker(
	su student.IStudentUsecase,
) student.IStudentWorker {
	return &studentWorker{
		su: su,
	}
}

func (sw *studentWorker) Start(kurl string) {
	krUserCreated := kafkaUtils.NewKafkaReader(kurl, initialize.KafkaAuthUserCreated, "auth-service")
	initialize.StartKafkaConsumer(krUserCreated, sw.HandleUserCreatedEvent)
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
