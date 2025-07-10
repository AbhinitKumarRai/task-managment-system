package service

import (
	"context"

	"github.com/AbhinitKumarRai/task-service/pkg/utils"
	taskPb "github.com/AbhinitKumarRai/task-service/proto"
)

type TaskGRPCServer struct {
	taskPb.UnimplementedTaskServiceServer
	svc *TaskService
}

func NewTaskGRPCServer(svc *TaskService) *TaskGRPCServer {
	return &TaskGRPCServer{svc: svc}
}

func (s *TaskGRPCServer) CreateTask(ctx context.Context, req *taskPb.CreateTaskRequest) (*taskPb.CreateTaskResponse, error) {
	task := utils.ProtoToModelTask(req.Task)
	err := s.svc.Create(task)
	if err != nil {
		return &taskPb.CreateTaskResponse{}, err
	}
	return &taskPb.CreateTaskResponse{}, nil
}

func (s *TaskGRPCServer) GetTask(ctx context.Context, req *taskPb.GetTaskRequest) (*taskPb.GetTaskResponse, error) {
	task, err := s.svc.GetByID(int(req.Id), int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &taskPb.GetTaskResponse{Task: utils.ModelToProtoTask(&task)}, nil
}

func (s *TaskGRPCServer) UpdateTask(ctx context.Context, req *taskPb.UpdateTaskRequest) (*taskPb.UpdateTaskResponse, error) {
	task := utils.ProtoToModelTask(req.UpdatedTask)
	err := s.svc.Update(*task)
	if err != nil {
		return nil, err
	}
	return &taskPb.UpdateTaskResponse{}, nil
}

func (s *TaskGRPCServer) DeleteTask(ctx context.Context, req *taskPb.DeleteTaskRequest) (*taskPb.DeleteTaskResponse, error) {
	err := s.svc.Delete(int(req.Id), int(req.UserId))
	if err != nil {
		return nil, err
	}

	return &taskPb.DeleteTaskResponse{}, nil
}

func (s *TaskGRPCServer) ListTasks(ctx context.Context, req *taskPb.ListTasksRequest) (*taskPb.ListTasksResponse, error) {
	tasks, err := s.svc.List(int(req.UserId), int(req.Page), int(req.Limit), utils.ProtoToModelStatus(req.Status), req.SortByDate)
	if err != nil {
		return nil, err
	}

	var protoTasks []*taskPb.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, utils.ModelToProtoTask(&task))
	}
	return &taskPb.ListTasksResponse{Tasks: protoTasks}, nil
}

func (s *TaskGRPCServer) DeleteAllTaskOfUser(ctx context.Context, req *taskPb.DeleteAllTaskOfUserRequest) (*taskPb.DeleteAllTaskOfUserResponse, error) {
	s.svc.DeleteAllTasksOfUser(int(req.UserId))

	return &taskPb.DeleteAllTaskOfUserResponse{}, nil
}
