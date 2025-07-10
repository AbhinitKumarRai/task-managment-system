package model

import "time"

type TaskStatus int

const (
	NoStatus TaskStatus = iota
	NotStarted
	Started
	Pending
	IsTriggered
	NotTriggered
	Completed
	InProgress
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserID      int        `json:"user_id"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	TriggerAt   time.Time  `json:"trigger_at,omitempty"`
	Triggered   bool       `json:"triggered"`
}
