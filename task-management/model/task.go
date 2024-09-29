package model

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}
