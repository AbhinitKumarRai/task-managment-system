package utils

import (
	"github.com/AbhinitKumarRai/task-service/pkg/model"
	taskPb "github.com/AbhinitKumarRai/task-service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtoToModelStatus(p taskPb.TaskStatus) model.TaskStatus {
	return model.TaskStatus(p) // Safe if enums align
}

func ModelToProtoStatus(s model.TaskStatus) taskPb.TaskStatus {
	return taskPb.TaskStatus(s) // Safe if enums align
}

func ProtoToModelTask(p *taskPb.Task) *model.Task {
	task := model.Task{
		ID:          int(p.Id),
		Title:       p.Title,
		Description: p.Description,
		UserID:      int(p.UserId),
		Status:      ProtoToModelStatus(p.Status),
		CreatedAt:   p.CreatedAt.AsTime(),
		UpdatedAt:   p.UpdatedAt.AsTime(),
		Triggered:   p.Triggered,
	}

	if p.TriggerAt != nil {
		task.TriggerAt = p.TriggerAt.AsTime()
	}

	return &task
}

func ModelToProtoTask(t *model.Task) *taskPb.Task {
	p := &taskPb.Task{
		Id:          int32(t.ID),
		Title:       t.Title,
		Description: t.Description,
		UserId:      int32(t.UserID),
		Status:      ModelToProtoStatus(t.Status),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
		Triggered:   t.Triggered,
		TriggerAt:   timestamppb.New(t.TriggerAt),
	}
	return p
}
