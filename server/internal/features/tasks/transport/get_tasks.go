package tasks_transport

import (
	"fmt"
	"net/http"

	core_logger "github.com/Fioneo/taskflow/internal/core/logger"
	core_http_response "github.com/Fioneo/taskflow/internal/core/transport/http/response"
	core_http_utils "github.com/Fioneo/taskflow/internal/core/transport/http/utils"
)

type GetTasksResponse []TaskDTOResponse

func (h *tasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	limit, offset, userID, err := GetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query params")
		return
	}

	taskDomains, err := h.tasksService.GetTasks(ctx, limit, offset, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}

	response := GetTasksResponse(tasksDTOFromDomains(taskDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}
func GetQueryParams(r *http.Request) (*int, *int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParams(r, "limit")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_utils.GetIntQueryParams(r, "offset")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	id, err := core_http_utils.GetIntQueryParams(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}
	return limit, offset, id, nil
}
