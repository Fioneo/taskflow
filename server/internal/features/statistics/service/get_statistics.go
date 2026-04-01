package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(ctx context.Context, id *int, from *time.Time, to *time.Time) (domain.Statistics, error) {
	if to != nil && from != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf("'to' must be after 'from': %w", core_errors.ErrInvalidArgument)
		}
	}
	tasks, err := s.StatisticsRepository.GetTasks(ctx, id, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("failed to get tasks: %w", err)
	}

	statistics := calcStatistics(tasks)
	return statistics, nil
}
func calcStatistics(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.Statistics{}
	}
	tasksCreated := len(tasks)
	tasksCompleted := 0
	var totalCompletedDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			tasksCompleted++
		}
		completionDuration := task.CompletionDuration()
		if completionDuration != nil {
			totalCompletedDuration += *completionDuration
		}
	}

	tasksCompletedRate := float64(tasksCompleted) / float64(tasksCreated) * 100
	var tasksACT *time.Duration

	if tasksCompleted > 0 && totalCompletedDuration != 0 {
		avg := totalCompletedDuration / time.Duration(tasksCompleted)

		tasksACT = &avg
	}

	return domain.Statistics{
		TasksCreated:       tasksCreated,
		TasksCompleted:     tasksCompleted,
		TasksCompletedRate: &tasksCompletedRate,
		TasksACT:           tasksACT,
	}
}
