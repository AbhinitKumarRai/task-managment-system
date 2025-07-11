package main

import (
	"log"
	"net"
	"os"

	"github.com/AbhinitKumarRai/task-service/internal/service"
	"github.com/AbhinitKumarRai/task-service/internal/taskmanager"
	kafkaPkg "github.com/AbhinitKumarRai/task-service/pkg/kafka"
	taskPb "github.com/AbhinitKumarRai/task-service/proto"
	"google.golang.org/grpc"
)

func main() {

	kafkaWriterInst, err := kafkaPkg.ConnectToKafkaAndCreateWriterInst()
	if err != nil {
		log.Panic(err)
	}

	taskManager := taskmanager.NewTaskManager()

	taskService := service.NewTaskService(taskManager, kafkaWriterInst)

	grpcServer := grpc.NewServer()
	taskPb.RegisterTaskServiceServer(grpcServer, service.NewTaskGRPCServer(taskService))

	port := os.Getenv("GRPC_PORT")
	listenAddr := ":" + port

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Task Service gRPC server running on %s", listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
