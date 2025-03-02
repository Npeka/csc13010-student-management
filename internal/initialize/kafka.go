package initialize

import (
	"context"
	"fmt"
	"log"

	"github.com/csc13010-student-management/config"
	kafkaUtils "github.com/csc13010-student-management/pkg/kafka"
	"github.com/segmentio/kafka-go"
)

const (
	KafkaAuthUserCreated     = "auth.user.created"
	KafkaAuthUserUpdated     = "auth.user.updated"
	KafkaStudentCreated      = "student.created"
	KafkaStudentUpdated      = "student.updated"
	KafkaNotificationCreated = "notification.created"
	KafkaNotificationUpdated = "notification.updated"
)

func NewKafkaTopics(cfg config.KafkaConfig) {
	kurl := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	topics := []string{
		KafkaAuthUserCreated,
		KafkaAuthUserUpdated,
		KafkaStudentCreated,
		KafkaStudentUpdated,
		KafkaNotificationCreated,
		KafkaNotificationUpdated,
	}

	for _, topic := range topics {
		kafkaUtils.NewKafkaTopic(kurl, topic)
	}
}

func StartKafkaConsumer(kafkaReader *kafka.Reader, handleFunc func(context.Context, kafka.Message) error) {
	ctx := context.Background()
	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		if err := handleFunc(ctx, msg); err != nil {
			log.Printf("Error handling message: %v", err)
		}
	}
}
