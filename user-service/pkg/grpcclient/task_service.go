package grpcclient

import (
	"context"

	taskPb "github.com/AbhinitKumarRai/user-service/proto"
	"google.golang.org/grpc"
)

type TaskGRPCClient struct {
	client taskPb.TaskServiceClient
}

func NewTaskGRPCClient(conn *grpc.ClientConn) *TaskGRPCClient {
	return &TaskGRPCClient{client: taskPb.NewTaskServiceClient(conn)}
}

func (c *TaskGRPCClient) CreateTask(ctx context.Context, req *taskPb.CreateTaskRequest) (*taskPb.CreateTaskResponse, error) {
	return c.client.CreateTask(ctx, req)
}

func (c *TaskGRPCClient) GetTask(ctx context.Context, req *taskPb.GetTaskRequest) (*taskPb.GetTaskResponse, error) {
	return c.client.GetTask(ctx, req)
}

func (c *TaskGRPCClient) UpdateTask(ctx context.Context, req *taskPb.UpdateTaskRequest) (*taskPb.UpdateTaskResponse, error) {
	return c.client.UpdateTask(ctx, req)
}

func (c *TaskGRPCClient) DeleteTask(ctx context.Context, req *taskPb.DeleteTaskRequest) (*taskPb.DeleteTaskResponse, error) {
	return c.client.DeleteTask(ctx, req)
}

func (c *TaskGRPCClient) DeleteAllTaskOfUser(ctx context.Context, req *taskPb.DeleteAllTaskOfUserRequest) (*taskPb.DeleteAllTaskOfUserResponse, error) {
	return c.client.DeleteAllTaskOfUser(ctx, req)
}

func (c *TaskGRPCClient) ListTasks(ctx context.Context, req *taskPb.ListTasksRequest) (*taskPb.ListTasksResponse, error) {
	return c.client.ListTasks(ctx, req)
}
