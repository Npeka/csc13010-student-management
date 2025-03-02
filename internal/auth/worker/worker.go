package worker

import (
	"context"
	"encoding/json"

	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/auth/dtos"
	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/initialize"
	kafkaUtils "github.com/csc13010-student-management/pkg/kafka"
	"github.com/segmentio/kafka-go"
)

type authWorker struct {
	au auth.IAuthUsecase
}

func NewAuthWorker(
	au auth.IAuthUsecase,
) auth.IAuthWorker {
	return &authWorker{
		au: au,
	}
}

func (aw *authWorker) Start(kurl string) {
	krStudentCreated := kafkaUtils.NewKafkaReader(kurl, initialize.KafkaStudentCreated, "student-service")
	initialize.StartKafkaConsumer(krStudentCreated, aw.HandleStudentCreatedEvent)
}

func (aw *authWorker) HandleStudentCreatedEvent(ctx context.Context, msg kafka.Message) error {
	var event events.StudentCreatedEvent
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	registerReq := &dtos.UserRegisterRequestDTO{
		Username: event.Username,
		Password: event.Username,
		Role:     event.Role,
	}

	if err := aw.au.Register(ctx, registerReq); err != nil {
		return err
	}

	return nil
}
