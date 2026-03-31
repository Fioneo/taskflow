package tasks_transport

import (
	"context"
	"net/http"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_http_server "github.com/Fioneo/taskflow/internal/core/transport/http/server"
)

type tasksHTTPHandler struct {
	tasksService TasksService
}
type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	GetTasks(ctx context.Context, limit, offset, userID *int) ([]domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	PatchTask(ctx context.Context, id int, patch domain.TaskPatch) (domain.Task, error)
}

func NewTasksHTTPHandler(tasksService TasksService) *tasksHTTPHandler {
	return &tasksHTTPHandler{
		tasksService: tasksService,
	}
}
func (h *tasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: h.PatchTask,
		},
	}
}
