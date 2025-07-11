package kafka

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func ConnectToKafkaAndCreateWriterInst() (*kafka.Writer, error) {
	// Kafka writer setup
	var kafkaW *kafka.Writer
	brokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")
	brokerList := strings.Split(brokers, ",")

	err := WaitForKafka(brokerList[0], 10, 3*time.Second)
	if err != nil {
		return nil, fmt.Errorf("kafka not ready: %v", err)
	}
	err = CreateTopic(brokerList, topic)
	if err != nil {
		return nil, fmt.Errorf("topic creation failed: %v", err)
	}
	if brokers != "" && topic != "" {
		kafkaW = &kafka.Writer{
			Addr:  kafka.TCP(brokerList...),
			Topic: topic,
		}
		log.Printf("Task Service: Kafka producer enabled for topic %s", topic)
	}

	return kafkaW, nil
}

func WaitForKafka(broker string, retries int, delay time.Duration) error {
	for i := 0; i < retries; i++ {
		conn, err := kafka.Dial("tcp", broker)
		if err == nil {
			conn.Close()
			log.Printf("Kafka is available at %s", broker)
			return nil
		}
		log.Printf("Kafka not ready yet (%d/%d): %v", i+1, retries, err)
		time.Sleep(delay)
	}
	return fmt.Errorf("kafka not available at %s after %d retries", broker, retries)
}

func CreateTopic(brokers []string, topic string) error {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		return fmt.Errorf("failed to dial kafka: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to get controller: %w", err)
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return fmt.Errorf("failed to dial controller: %w", err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	log.Printf("Kafka topic %s created", topic)
	return nil
}
