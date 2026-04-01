package domain

import "time"

type Statistics struct {
	TasksCreated       int
	TasksCompleted     int
	TasksCompletedRate *float64
	TasksACT           *time.Duration
}

// TaskACT = TaskAverageCompletionTime

func NewStatistics(
	tasksCreated int,
	taskCompleted int,
	tasksCompletedRate *float64,
	taskACT *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:       tasksCreated,
		TasksCompleted:     taskCompleted,
		TasksCompletedRate: tasksCompletedRate,
		TasksACT:           taskACT,
	}
}
