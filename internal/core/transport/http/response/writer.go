package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter

	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitialized,
	}
}

// Добавляем методу WriteHeader сохранение указанного статус кода
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// Получение сохранённого статус кода
func (w *ResponseWriter) GetStatusCode() int {
	if w.statusCode == StatusCodeUninitialized {
		return http.StatusOK
	}

	return w.statusCode
}
