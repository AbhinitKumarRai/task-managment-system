package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	tm "github.com/AbhinitKumarRai/task-service/internal/taskmanager"
	"github.com/AbhinitKumarRai/task-service/pkg/model"
	"github.com/segmentio/kafka-go"
)

type TaskService struct {
	taskManager *tm.TaskManager
	kafkaW      *kafka.Writer
}

func NewTaskService(taskManager *tm.TaskManager, kafkaW *kafka.Writer) *TaskService {
	ts := &TaskService{
		taskManager: taskManager,
		kafkaW:      kafkaW,
	}

	go ts.startScheduler()
	return ts
}

func (s *TaskService) startScheduler() {
	for {
		now := time.Now()
		dueTasks := s.taskManager.ListScheduledDue(now)
		for _, t := range dueTasks {

			message := fmt.Sprintf("Reminder for your todo list item with name %s", t.Title)
			err := s.SendToKafka(t.ID, t.UserID, "scheduled_task_reminder", message, "random@random", "0123456789")
			if err != nil {
				fmt.Println(fmt.Errorf("unable to send to kafka: %v", err))
			} else {
				log.Printf("[Scheduler] Triggered notification for task %d at %s", t.ID, now.Format(time.RFC3339))
			}

		}
		time.Sleep(time.Minute)
	}
}

func (s *TaskService) Create(task *model.Task) (model.Task, error) {
	createdTask, err := s.taskManager.Create(task)
	if err != nil {
		return createdTask, err
	}

	kafkaMsg := fmt.Sprintf("We have successfully created your task with title: %s and id: %d", createdTask.Title, createdTask.ID)
	err = s.SendToKafka(createdTask.ID, createdTask.UserID, "task_created", kafkaMsg, "random@random.com", "123456789")
	if err != nil {
		return model.Task{}, err
	}

	return createdTask, nil
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

func (s *TaskService) DeleteAllTasksOfUser(userId int) error {
	return s.taskManager.DeleteAllTaskOfUser(userId)
}

func (s *TaskService) List(userId, page, limit int, status model.TaskStatus, sortByDueDate bool) ([]model.Task, error) {
	return s.taskManager.ListTasksOfUser(userId, page, limit, status, sortByDueDate)
}

func (s *TaskService) SendToKafka(taskId, userId int, taskType, message, email, number string) error {
	if s.kafkaW == nil {
		return fmt.Errorf("no kafka instance initialized")
	}

	event := map[string]interface{}{
		"task_id": taskId,
		"user_id": userId,
		"type":    taskType,
		"message": message,
		"email":   email,
		"number":  number,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal kafka event: %w", err)
	}

	err = s.kafkaW.WriteMessages(context.Background(), kafka.Message{Value: data})
	if err != nil {
		return fmt.Errorf("failed to write kafka message: %w", err)
	}

	log.Printf("[Kafka] Published message for task %d: %s", taskId, string(data))
	return nil
}
