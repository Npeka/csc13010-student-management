package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/auth/dtos"
	"github.com/csc13010-student-management/internal/initialize"
	kafkaUtils "github.com/csc13010-student-management/pkg/kafka"
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
	// Tạo nhiều consumer cùng group
	topic := "dbserver1.public.students"
	groupID := "auth-service" // Consumer group chung
	initialize.WaitForTopic(kurl, topic, 5*time.Minute)
	kr1 := kafkaUtils.NewKafkaReader(kurl, topic, groupID)
	kr2 := kafkaUtils.NewKafkaReader(kurl, topic, groupID)

	// Chạy song song
	go initialize.StartKafkaConsumer(kr1, aw.HandleStudentCreatedEvent)
	go initialize.StartKafkaConsumer(kr2, aw.HandleStudentCreatedEvent)
}

func (aw *authWorker) HandleStudentCreatedEvent(ctx context.Context, msg kafka.Message) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authWorker.HandleStudentCreatedEvent")
	defer span.Finish()

	var data = make(map[string]interface{})
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		logger.LoggerFuncError(aw.lg, errors.Wrap(err, "authWorker.HandleStudentCreatedEvent.json.Unmarshal"))
		return err
	}

	payload, ok := data["payload"].(map[string]interface{})
	if !ok {
		return nil
	}

	before, ok := payload["before"].(map[string]interface{})
	if ok && len(before) > 0 {
		return nil
	}

	after, ok := payload["after"].(map[string]interface{})

	if ok {
		studentID, ok := after["student_id"].(string)
		if !ok || studentID == "" {
			return nil
		}

		registerReq := &dtos.UserRegisterRequestDTO{
			Username: studentID,
			Password: studentID,
			Role:     "student",
		}
		if err := aw.au.Register(ctx, registerReq); err != nil {
			logger.LoggerFuncError(aw.lg, errors.Wrap(err, "authWorker.HandleStudentCreatedEvent.au.Register"))
			return err
		}
	}

	return nil
}

// var event events.StudentCreatedEvent
// if err := json.Unmarshal(msg.Value, &event); err != nil {
// 	return err
// }

// registerReq := &dtos.UserRegisterRequestDTO{
// 	Username: event.Username,
// 	Password: event.Username,
// 	Role:     event.Role,
// }

// if err := aw.au.Register(ctx, registerReq); err != nil {
// 	return err
// }
