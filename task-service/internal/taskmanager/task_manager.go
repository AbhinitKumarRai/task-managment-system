package taskmanager

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/AbhinitKumarRai/task-service/pkg/model"
)

type TaskManager struct {
	globalLock     *sync.Mutex
	tasksByUserId  map[int]map[int]*model.Task
	userLevelLocks map[int]*sync.RWMutex
	taskIDCounter  int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		globalLock:     &sync.Mutex{},
		tasksByUserId:  make(map[int]map[int]*model.Task),
		userLevelLocks: make(map[int]*sync.RWMutex),
		taskIDCounter:  1,
	}
}

func (m *TaskManager) getUserLock(userID int) *sync.RWMutex {
	m.globalLock.Lock()
	defer m.globalLock.Unlock()

	lock, exists := m.userLevelLocks[userID]
	if !exists {
		lock = &sync.RWMutex{}
		m.userLevelLocks[userID] = lock
	}
	return lock
}

func (m *TaskManager) Create(task *model.Task) (model.Task, error) {
	userlock := m.getUserLock(task.UserID)
	userlock.Lock()
	defer userlock.Unlock()

	m.globalLock.Lock()
	task.ID = m.taskIDCounter
	m.taskIDCounter++
	m.globalLock.Unlock()

	if _, ok := m.tasksByUserId[task.UserID]; !ok {
		m.tasksByUserId[task.UserID] = make(map[int]*model.Task)
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	m.tasksByUserId[task.UserID][task.ID] = task

	return *task, nil
}

func (m *TaskManager) GetTaskByID(id, userId int) (model.Task, error) {
	userlock := m.getUserLock(userId)
	userlock.RLock()
	defer userlock.RUnlock()

	if taskList, ok := m.tasksByUserId[userId]; ok {
		if task, ok := taskList[id]; ok {
			return *task, nil
		}

		return model.Task{}, fmt.Errorf("task with id %d doesnot belong to user %d", id, userId)
	}
	return model.Task{}, fmt.Errorf("user not found for id: %d", userId)
}

func (m *TaskManager) Update(update model.Task) error {
	userlock := m.getUserLock(update.UserID)
	userlock.Lock()
	defer userlock.Unlock()

	if update.ID == 0 || update.UserID == 0 {
		return fmt.Errorf("user id or task id cannot be empty")
	}
	if taskList, ok := m.tasksByUserId[update.UserID]; ok {
		if task, ok := taskList[update.ID]; ok {
			if update.Description != "" {
				task.Description = update.Description
			}

			if update.Status != model.NoStatus {
				task.Status = update.Status
			}

			if update.TriggerAt.IsZero() {
				task.TriggerAt = update.TriggerAt
			}

			task.UpdatedAt = time.Now()
			m.tasksByUserId[update.UserID][update.ID] = task

			return nil
		}

		return fmt.Errorf("task with id %d doesnot belong to user %d", update.ID, update.UserID)
	}
	return fmt.Errorf("user not found for id: %d", update.UserID)
}

func (m *TaskManager) DeleteTask(id, userId int) error {
	userlock := m.getUserLock(userId)
	userlock.Lock()
	defer userlock.Unlock()
	if taskList, ok := m.tasksByUserId[userId]; ok {
		if _, ok := taskList[id]; ok {
			delete(m.tasksByUserId[userId], id)

		}
		return fmt.Errorf("task with id %d doesnot belong to user %d", id, userId)
	}

	return fmt.Errorf("user not found for id: %d", userId)
}

func (m *TaskManager) DeleteAllTaskOfUser(userId int) error {
	m.globalLock.Lock()
	defer m.globalLock.Unlock()

	delete(m.tasksByUserId, userId)

	delete(m.userLevelLocks, userId)

	return nil
}

func (m *TaskManager) ListTasksOfUser(userId, page, limit int, status model.TaskStatus, sortByDate bool) ([]model.Task, error) {
	userlock := m.getUserLock(userId)
	userlock.RLock()
	defer userlock.RUnlock()
	filtered := []model.Task{}
	if taskList, ok := m.tasksByUserId[userId]; ok {
		for _, task := range taskList {
			if status == model.NoStatus || task.Status == status {
				filtered = append(filtered, *task)
			}
		}
		if sortByDate {
			sort.Slice(filtered, func(i, j int) bool {
				return filtered[i].CreatedAt.Before(filtered[j].CreatedAt)
			})
		}
		start := (page - 1) * limit
		end := start + limit
		if start > len(filtered) {
			start = len(filtered)
		}
		if end > len(filtered) {
			end = len(filtered)
		}
		return filtered[start:end], nil
	}

	return filtered, fmt.Errorf("user not found for id: %d", userId)
}

// ListScheduledDue returns tasks with TriggerAt <= now and not Triggered
func (m *TaskManager) ListScheduledDue(expectedTime time.Time) []model.Task {

	var due []model.Task
	for userId, taskList := range m.tasksByUserId {
		userlock := m.getUserLock(userId)
		userlock.RLock()
		defer userlock.RUnlock()
		for _, task := range taskList {
			if !task.TriggerAt.IsZero() && !task.Triggered && task.TriggerAt.Before(expectedTime.Add(1*time.Second)) {
				due = append(due, *task)
			}
		}
	}

	return due
}

// MarkTriggered sets Triggered=true for a task
func (m *TaskManager) MarkTriggered(id, userId int) {
	userlock := m.getUserLock(userId)
	userlock.Lock()
	defer userlock.Unlock()
	if taskList, ok := m.tasksByUserId[userId]; ok {
		if _, ok := taskList[id]; ok {
			m.tasksByUserId[userId][id].Triggered = true
		}
	}
}
