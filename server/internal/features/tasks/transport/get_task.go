package tasks_transport

import (
	"net/http"

	core_logger "github.com/Fioneo/taskflow/internal/core/logger"
	core_http_response "github.com/Fioneo/taskflow/internal/core/transport/http/response"
	core_http_utils "github.com/Fioneo/taskflow/internal/core/transport/http/utils"
)

type GetTaskResponse TaskDTOResponse

func (h *tasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	id, err := core_http_utils.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'id' from path value")
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
