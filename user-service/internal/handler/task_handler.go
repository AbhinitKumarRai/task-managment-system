package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/AbhinitKumarRai/user-service/pkg/grpcclient"
	"github.com/AbhinitKumarRai/user-service/pkg/middleware"
	taskPb "github.com/AbhinitKumarRai/user-service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskHandler struct {
	grpcClient *grpcclient.TaskGRPCClient
}

func NewTaskHandler(grpcClient *grpcclient.TaskGRPCClient) *TaskHandler {
	return &TaskHandler{grpcClient: grpcClient}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		TriggerAt   *string `json:"trigger_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var triggerAt *timestamppb.Timestamp
	if req.TriggerAt != nil && *req.TriggerAt != "" {
		t, err := time.Parse(time.RFC3339, *req.TriggerAt)
		if err != nil {
			http.Error(w, "invalid trigger_at format", http.StatusBadRequest)
			return
		}
		triggerAt = timestamppb.New(t)
	}

	task := &taskPb.Task{
		Title:       req.Title,
		Description: req.Description,
		UserId:      int32(userID),
		Status:      taskPb.TaskStatus(0),
		TriggerAt:   triggerAt,
	}

	resp, err := h.grpcClient.CreateTask(context.Background(), &taskPb.CreateTaskRequest{Task: task})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.TaskCreated)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var req struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Status      string  `json:"status"`
		TriggerAt   *string `json:"trigger_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)

	var triggerAt *timestamppb.Timestamp
	if req.TriggerAt != nil && *req.TriggerAt != "" {
		t, err := time.Parse(time.RFC3339, *req.TriggerAt)
		if err != nil {
			http.Error(w, "invalid trigger_at format", http.StatusBadRequest)
			return
		}
		triggerAt = timestamppb.New(t)
	}

	var statusEnum int32

	if req.Status == "" {
		statusEnum = 0
	}

	statusEnum, ok := taskPb.TaskStatus_value[req.Status]
	if !ok {
		http.Error(w, "invalid status", http.StatusBadRequest)
		return
	}

	task := &taskPb.Task{
		Id:          int32(id),
		Title:       req.Title,
		Description: req.Description,
		UserId:      int32(userID),
		Status:      taskPb.TaskStatus(statusEnum),
		TriggerAt:   triggerAt,
	}

	_, err := h.grpcClient.UpdateTask(context.Background(), &taskPb.UpdateTaskRequest{
		UpdatedTask: task,
	})
	if err != nil {
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)

	statusStr := r.URL.Query().Get("status")
	status := taskPb.TaskStatus_TASK_STATUS_NO_STATUS
	if statusStr != "" {
		if val, ok := taskPb.TaskStatus_value[statusStr]; ok {
			status = taskPb.TaskStatus(val)
		} else {
			http.Error(w, "invalid status", http.StatusBadRequest)
			return
		}
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}
	sortByDate := r.URL.Query().Get("sort") == "date"

	req := &taskPb.ListTasksRequest{
		UserId:     int32(userID),
		Page:       int32(page),
		Limit:      int32(limit),
		Status:     status,
		SortByDate: sortByDate,
	}
	resp, err := h.grpcClient.ListTasks(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	userID := r.Context().Value(middleware.UserIDKey).(int)

	resp, err := h.grpcClient.GetTask(context.Background(), &taskPb.GetTaskRequest{
		Id:     int32(id),
		UserId: int32(userID),
	})
	if err != nil || resp.Task == nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Task)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	userID := r.Context().Value(middleware.UserIDKey).(int)

	_, err := h.grpcClient.DeleteTask(context.Background(), &taskPb.DeleteTaskRequest{
		Id:     int32(id),
		UserId: int32(userID),
	})
	if err != nil {
		http.Error(w, "delete failed", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
