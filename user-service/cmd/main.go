package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbhinitKumarRai/user-service/internal/routes"
	"github.com/AbhinitKumarRai/user-service/internal/service"
	manager "github.com/AbhinitKumarRai/user-service/internal/usermanager"
	"github.com/AbhinitKumarRai/user-service/pkg/grpcclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	taskServiceAddr := os.Getenv("TASK_SERVICE_ADDR")
	if taskServiceAddr == "" {
		taskServiceAddr = "localhost:50051"
	}
	conn := connectToTaskService(taskServiceAddr)
	defer conn.Close()

	userManager := manager.NewUserManager()
	grpcClient := grpcclient.NewTaskGRPCClient(conn)
	userService := service.NewUserService(userManager)

	// Register routes using Gorilla Mux
	router := routes.RegisterRoutes(userService, grpcClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("User Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func connectToTaskService(addr string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // blocks until connection is established or timeout
	)
	if err != nil {
		log.Fatalf("failed to connect to task service at %s: %v", addr, err)
	}
	log.Printf("Connected to task service at %s", addr)
	return conn
}
