package statistics_transport

import (
	"context"
	"net/http"
	"time"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_http_server "github.com/Fioneo/taskflow/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	StatisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(ctx context.Context, userID *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		StatisticsService: statisticsService,
	}
}
func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
