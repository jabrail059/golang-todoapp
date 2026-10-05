package core_http_middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	core_logger "github.com/jabrail059/golang-todoapp/internal/core/logger"
	core_http_response "github.com/jabrail059/golang-todoapp/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
)

// CORS-обработчик проверяет, входит ли сайт с входящим запросом в список доверенных адресов
func CORS(allowedOriginsList []string) Middleware {
	// Множество доверенных адресов
	allowedOrigins := map[string]struct{}{}
	for _, origin := range allowedOriginsList {
		allowedOrigins[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем адрес сайта и проверяем, есть ли он в множестве или нет
			// Если входит, то проставляем соответствующие заголовки
			origin := r.Header.Get("Origin")
			if _, ok := allowedOrigins[origin]; ok {
				// Сообщаем браузеру, что запросы с данного сайта разрешены
				w.Header().Set("Access-Control-Allow-Origin", origin)
				// Прописываем разрешенные методы для него
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
				// Прописываем разрешенные http-заголовки запроса
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			// Проверяем метод запроса:
			// если метод Options, значит браузер проверяет, разрешен ли запрос
			// нужно вернуть код 200
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			// В противном случае продолжаем обработку http-запроса
			next.ServeHTTP(w, r)
		})
	}
}

// Получаем id каждого HTTP запроса для удобного отслеживания в логах
func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

// Добавляем логгер в контекст каждого HTTP запроса, тем самым "пробрасывая" его для дальнейшего использования
func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := core_logger.ToContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Отслеживаем всё, что происходит с HTTP запросом, до конца его выполнения
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()

			log.Debug(
				">>> incoming HTTP request",
				zap.String("http_method", r.Method),
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(before)),
			)
		})
	}
}

// Корректная обработка паники при обработке HTTP запроса без падения сервера
func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
