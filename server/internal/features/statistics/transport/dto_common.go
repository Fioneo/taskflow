package statistics_transport

import "github.com/Fioneo/taskflow/internal/core/domain"

type StatisticsDTOResponse struct {
	TasksCreated       int      `json:"tasks_created"`
	TasksCompleted     int      `json:"tasks_completed"`
	TasksCompletedRate *float64 `json:"tasks_completed_rate"`
	TasksACT           *string  `json:"tasks_average_completion_time"`
}

func toDTOFromDomain(statistics domain.Statistics) StatisticsDTOResponse {
	var avgTime *string
	if statistics.TasksACT != nil {
		duration := statistics.TasksACT.String()
		avgTime = &duration
	}
	return StatisticsDTOResponse{
		TasksCreated:       statistics.TasksCreated,
		TasksCompleted:     statistics.TasksCompleted,
		TasksCompletedRate: statistics.TasksCompletedRate,
		TasksACT:           avgTime,
	}
}
