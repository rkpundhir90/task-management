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

type CreateTaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id" binding:"required"`
	Status      string    `json:"status" binding:"required,oneof=Pending In Progress Completed"`
	DueDate     time.Time `json:"due_date" binding:"required"`
}

type UpdateTaskRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required,oneof=Pending In Progress Completed"`
}
