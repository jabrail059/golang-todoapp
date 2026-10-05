package tasks_transport_http

import (
	"time"

	"github.com/jabrail059/golang-todoapp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"             example:"95"`
	Version      int        `json:"version"        example:"6"`
	Title        string     `json:"title"          example:"Домашка"`
	Description  *string    `json:"description"    example:"Сделать дз по физре до завтра"`
	Completed    bool       `json:"completed"      example:"false"`
	CreatedAt    time.Time  `json:"created_at"     example:"2026-04-10T22:34:05Z"`
	CompletedAt  *time.Time `json:"completed_at"   example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"5"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	tasksDTO := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		tasksDTO[i] = taskDTOFromDomain(task)
	}

	return tasksDTO
}
