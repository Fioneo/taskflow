package users_transport

import (
	"net/http"

	core_logger "github.com/Fioneo/taskflow/internal/core/logger"
	core_http_response "github.com/Fioneo/taskflow/internal/core/transport/http/response"
	core_http_utils "github.com/Fioneo/taskflow/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	id, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id from path value")
		return
	}
	userDomain, err := h.usersService.GetUser(ctx, id)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}
	response := GetUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
