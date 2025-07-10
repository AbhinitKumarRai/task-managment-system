package main

import (
	"log"
	"net"

	"github.com/AbhinitKumarRai/task-service/internal/service"
	"github.com/AbhinitKumarRai/task-service/internal/taskmanager"
	taskPb "github.com/AbhinitKumarRai/task-service/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load()

	taskManager := taskmanager.NewTaskManager()

	taskService := service.NewTaskService(taskManager)

	grpcServer := grpc.NewServer()
	taskPb.RegisterTaskServiceServer(grpcServer, service.NewTaskGRPCServer(taskService))

	listenAddr := ":50051"
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Task Service gRPC server running on %s", listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
