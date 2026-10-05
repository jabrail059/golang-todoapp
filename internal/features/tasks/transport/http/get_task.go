package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/golang-todoapp/internal/core/logger"
	core_http_request "github.com/jabrail059/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/jabrail059/golang-todoapp/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

// GetTask 		godoc
// @Summary 	Получение задачи
// @Description Получение существующей в системе задачи по его ID
// @Tags 		tasks
// @Produce 	json
// @Param 		id 			path int true 								"ID получаемой задачи"
// @Success 	200 		{object} GetTaskResponse 					"Задача успешно найдена"
// @Failure 	400 		{object} core_http_response.ErrorResponse 	"Bad Request"
// @Failure 	404 		{object} core_http_response.ErrorResponse 	"Task Not Found"
// @Failure 	500 		{object} core_http_response.ErrorResponse 	"Internal Server Error"
// @Router 		/tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	taskDomain, err := h.TasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task",
		)

		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
