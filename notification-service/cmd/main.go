package main

import (
	"log"
	"os"
	"strings"

	"github.com/AbhinitKumarRai/notification-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	svc := service.NewNotificationService()

	// Optionally start Kafka consumer if env set
	brokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")
	if brokers != "" && topic != "" {
		log.Printf("Starting Kafka consumer for topic %s", topic)
		brokerList := strings.Split(brokers, ",")
		svc.ListenKafka(brokerList, topic)
	}

	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	log.Printf("Notification Service running on :%s", port)
	r.Run(":" + port)
}
