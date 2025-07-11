package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/AbhinitKumarRai/notification-service/pkg/model"
	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendMessage(notification model.Notification) error {
	fmt.Println("Sent Message to mobile==============> \n", notification)
	return nil
}

func (s *NotificationService) SendEmail(notification model.Notification) error {
	fmt.Println("Sent Message to email=========> \n", notification)
	return nil
}

// ListenKafka starts a Kafka consumer loop (mocked for now)
func (s *NotificationService) ListenKafka(brokers []string, topic string) {
	kafkaR := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "notification-service",
	})
	go func(kafkaR *kafka.Reader) {
		for {
			m, err := kafkaR.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Kafka error: %v", err)
				continue
			}
			log.Printf("[MOCK] Received Kafka message: %s", string(m.Value))

			var notification model.Notification

			err = json.Unmarshal(m.Value, &notification)
			if err != nil {
				log.Printf("data is not correct")
			}

			if notification.Email != "" {
				s.SendEmail(notification)
			}

			if notification.Number != "" {
				s.SendMessage(notification)
			}
		}
	}(kafkaR)
}
