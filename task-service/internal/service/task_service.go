package service

import (
	tm "github.com/AbhinitKumarRai/task-service/internal/taskmanager"
	"github.com/AbhinitKumarRai/task-service/pkg/model"
)

type TaskService struct {
	taskManager *tm.TaskManager
}

func NewTaskService(taskManager *tm.TaskManager) *TaskService {
	ts := &TaskService{taskManager: taskManager}
	return ts
}

func (s *TaskService) Create(task *model.Task) error {
	err := s.taskManager.Create(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetByID(id, userId int) (model.Task, error) {
	return s.taskManager.GetTaskByID(id, userId)
}

func (s *TaskService) Update(update model.Task) error {
	return s.taskManager.Update(update)
}

func (s *TaskService) Delete(id, userId int) error {
	return s.taskManager.DeleteTask(id, userId)
}

func (s *TaskService) List(userId, page, limit int, status model.TaskStatus, sortByDueDate bool) ([]model.Task, error) {
	return s.taskManager.ListTasksOfUser(userId, page, limit, status, sortByDueDate)
}
