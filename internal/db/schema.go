package db

import (
	"time"
)

type Status string
type Priority string

const (
	StatusPending 	 Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
	StatusDelayed    Status = "delayed"
)

const (
	LowPriority    Priority    = "low"
	MediumPriority Priority = "medium"
	HighPriority   Priority = "high"
)

type Job struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	Priority     Priority  `json:"priority"`
	MaxAttempts  int       `json:"max_attempts"`
	JobStatus    Status    `json:"job_status"`
	CurrAttempts int       `json:"curr_attempts"`
	Payload      string    `json:"payload"`
	CreatedAt    time.Time `json:"created_at"`
}

type JobCreation struct {
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Priority    Priority `json:"priority"`
	Payload     string   `json:"payload"`
	JobStatus   Status   `json:"job_status"`
}
