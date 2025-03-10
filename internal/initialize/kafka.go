package initialize

import (
	"fmt"

	"github.com/csc13010-student-management/config"
	"github.com/csc13010-student-management/internal/events"
	kafkas "github.com/csc13010-student-management/pkg/kafka"
)

const ()

func NewKafkaTopics(cfg config.KafkaConfig) {
	kurl := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	topics := []string{
		events.AuthUserCreated,
		events.AuthCreateUser,
		string(events.NotiStudentStatusChanged),
		string(events.NotiUserResetPasswordOTP),
	}

	for _, topic := range topics {
		kafkas.NewKafkaTopic(kurl, topic)
	}
}
