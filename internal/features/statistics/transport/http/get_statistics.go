package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/jabrail059/golang-todoapp/internal/core/logger"
	core_http_request "github.com/jabrail059/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/jabrail059/golang-todoapp/internal/core/transport/http/response"
)

type GetStatisticsResponse StatisticsDTOResponse

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	queryParams, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed get 'userID/from/to' query params",
		)

		return
	}

	statisticsDomain, err := h.statisticsService.GetStatistics(
		ctx,
		queryParams.userID,
		queryParams.from,
		queryParams.to,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get staticstics",
		)

		return
	}

	response := GetStatisticsResponse(toDTOFromDomain(statisticsDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

type queryParams struct {
	userID *int
	from   *time.Time
	to     *time.Time
}

func getUserIDFromToQueryParams(r *http.Request) (*queryParams, error) {
	const (
		toQueryParamKey     = "to"
		fromQueryParamKey   = "from"
		userIDQueryParamKey = "user_id"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetTimeQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetTimeQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return &queryParams{
		userID: userID,
		from:   from,
		to:     to,
	}, nil
}
