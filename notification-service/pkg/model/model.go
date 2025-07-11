package model

type Notification struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	TaskID  int    `json:"task_id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Email   string `json:"email"`
	Number  string `json:"number"`
}
