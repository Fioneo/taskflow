package tasks_transport

import (
	"fmt"
	"net/http"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_logger "github.com/Fioneo/taskflow/internal/core/logger"
	core_http_request "github.com/Fioneo/taskflow/internal/core/transport/http/request"
	core_http_response "github.com/Fioneo/taskflow/internal/core/transport/http/response"
	core_http_types "github.com/Fioneo/taskflow/internal/core/transport/http/types"
	core_http_utils "github.com/Fioneo/taskflow/internal/core/transport/http/utils"
)

type TaskPatchRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `jsob:"completed"`
}
type TaskPatchResponse TaskDTOResponse

func (r *TaskPatchRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("title can't be null")
		}
		titleLength := len([]rune(*r.Title.Value))
		if titleLength < 1 || titleLength > 30 {
			return fmt.Errorf("'Title' must be between 1 and 30 symbols")
		}
	}
	if r.Description.Set {
		if r.Description.Value != nil {
			DescriptionLength := len([]rune(*r.Description.Value))
			if DescriptionLength < 1 || DescriptionLength > 100 {
				return fmt.Errorf("'Description' must be between 1 and 100 symbols")
			}
		}
	}
	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("'Completed' can't be null")
		}
	}
	return nil
}
func (h *tasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	id, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'id' from path value")
		return
	}
	var request TaskPatchRequest

	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskPatch := TaskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, id, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch task")
		return
	}

	response := TaskPatchResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
func TaskPatchFromRequest(request TaskPatchRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
