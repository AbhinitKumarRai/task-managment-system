package main

// import (
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/AbhinitKumarRai/task-service/internal/handler"
// 	"github.com/AbhinitKumarRai/task-service/internal/repository"
// 	"github.com/AbhinitKumarRai/task-service/internal/service"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	"github.com/segmentio/kafka-go"
// )

// func main() {
// 	_ = godotenv.Load()

// 	repo := repository.NewTaskRepository()

// 	// Kafka writer setup
// 	var kafkaW *kafka.Writer
// 	brokers := os.Getenv("KAFKA_BROKERS")
// 	topic := os.Getenv("KAFKA_TOPIC")
// 	if brokers != "" && topic != "" {
// 		kafkaW = &kafka.Writer{
// 			Addr:  kafka.TCP(strings.Split(brokers, ",")...),
// 			Topic: topic,
// 		}
// 		log.Printf("Task Service: Kafka producer enabled for topic %s", topic)
// 	}

// 	svc := service.NewTaskService(repo, kafkaW)
// 	handler := handler.NewTaskHandler(svc)

// 	r := gin.Default()
// 	handler.RegisterRoutes(r)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8081"
// 	}
// 	log.Printf("Task Service running on :%s", port)
// 	r.Run(":" + port)
// }
