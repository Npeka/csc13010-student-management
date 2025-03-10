package kafkas

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	Topic   string
	GroupID string
	Handler func(context.Context, kafka.Message) error
	MaxIns  int
}

func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Consumer
func NewKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

func NewKafkaTopic(kafkaURL, topic string) {
	conn, err := kafka.Dial("tcp", kafkaURL)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}}

	// check if topic exists
	topics, err := controllerConn.ReadPartitions(topic)
	if err == nil && len(topics) > 0 {
		return
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func WaitForTopic(kurl, topic string, timeout time.Duration) error {
	start := time.Now()

	// Connect to Kafka
	conn, err := kafka.Dial("tcp", kurl)
	if err != nil {
		return fmt.Errorf("unable to connect to Kafka: %v", err)
	}
	defer conn.Close()

	for {
		// Get the list of topics from Kafka
		partitions, err := conn.ReadPartitions()
		if err != nil {
			return fmt.Errorf("error reading partitions: %v", err)
		}

		// Check if the topic exists
		for _, p := range partitions {
			if p.Topic == topic {
				log.Printf("Topic '%s' exists!", topic)
				return nil
			}
		}

		// If timeout is reached, exit the loop
		if time.Since(start) > timeout {
			return fmt.Errorf("timeout: topic '%s' was not created after %v", topic, timeout)
		}

		log.Println("Topic not ready, waiting for 2s...")
		time.Sleep(2 * time.Second) // Wait before checking again
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

func StartKafkaConsumers(kurl string, krs []KafkaReader) {
	var wg sync.WaitGroup
	for _, kr := range krs {
		for i := 0; i < kr.MaxIns; i++ {
			wg.Add(1)
			krc := kr
			go func(kr KafkaReader) {
				defer wg.Done()
				krnew := NewKafkaReader(kurl, kr.Topic, kr.GroupID)
				WaitForTopic(kurl, kr.Topic, 5*time.Minute)
				StartKafkaConsumer(krnew, kr.Handler)
			}(krc)
		}
	}
	wg.Wait()
}
