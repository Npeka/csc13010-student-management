package worker

import (
	"context"
	"encoding/json"

	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/auth/dtos"
	"github.com/csc13010-student-management/internal/events"
	kafkas "github.com/csc13010-student-management/pkg/kafka"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type authWorker struct {
	au auth.IAuthUsecase
	lg *logger.LoggerZap
}

func NewAuthWorker(
	au auth.IAuthUsecase,
	lg *logger.LoggerZap,
) auth.IAuthWorker {
	return &authWorker{
		au: au,
		lg: lg,
	}
}

func (aw *authWorker) Start(kurl string) {
	krs := []kafkas.KafkaReader{
		{
			Topic:   events.AuthCreateUser,
			GroupID: "auth.create.user",
			Handler: aw.HandleCreateUserEvent,
			MaxIns:  1,
		},
	}
	kafkas.StartKafkaConsumers(kurl, krs)
}

func (aw *authWorker) HandleCreateUserEvent(ctx context.Context, msg kafka.Message) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aw.HandleUserCreatedEvent")
	defer span.Finish()

	var data events.CreateUserEvent
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		logger.LoggerFuncError(aw.lg, errors.Wrap(err, "aw.HandleUserCreatedEvent.json.Unmarshal"))
		return err
	}

	registerReq := &dtos.UserRegisterRequestDTO{
		Username: data.Username,
		Password: data.Password,
		Role:     data.Role,
	}

	if err := aw.au.Register(ctx, registerReq); err != nil {
		logger.LoggerFuncError(aw.lg, errors.Wrap(err, "authWorker.HandleUserCreatedEvent.au.Register"))
		return err
	}

	return nil
}
