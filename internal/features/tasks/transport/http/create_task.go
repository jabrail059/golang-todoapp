package tasks_transport_http

import (
	"net/http"

	"github.com/jabrail059/golang-todoapp/internal/core/domain"
	core_logger "github.com/jabrail059/golang-todoapp/internal/core/logger"
	core_http_request "github.com/jabrail059/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/jabrail059/golang-todoapp/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"         example:"Домашнее задание"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000" example:"Сделать дз по матану до пятницы"`
	AuthorUserID int     `json:"author_user_id" validate:"required"              example:"5"`
}

type CreateTaskResponse TaskDTOResponse

// CreateTask 	godoc
// @Summary 	"Создание задачи"
// @Description "Создание новой задачи в системе"
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		request body 	 CreateTaskRequest 		true 		"CreateTask тело запроса"
// @Success 	201 	{object} CreateTaskResponse 				"Успешно созданная задача"
// @Failure 	400 	{object} core_http_response.ErrorResponse 	"Bad Request"
// @Failure 	404 	{object} core_http_response.ErrorResponse 	"Author Not Found"
// @Failure 	500 	{object} core_http_response.ErrorResponse 	"Internal Server Error"
// @Router 		/tasks 	[post]
func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke CreateTask handler")

	var request CreateTaskRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	taskDomain := domainFromDTO(request)

	taskDomain, err := h.TasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create task",
		)

		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(request CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)
}
