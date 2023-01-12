package util

import (
	"encoding/json"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b, err := json.Marshal(ErrorResponse{
		Timestamp: time.Now(),
		Status:    statusCode,
		Error:     http.StatusText(statusCode),
		Message:   err.Error(),
	})
	if err != nil {
		slog.Error("could not marshal error response", err)
	} else {
		_, err = w.Write(b)
		if err != nil {
			slog.Error("could not write error response", err)
		}
	}
}

type ApiHandlerFunc[T any] func(w http.ResponseWriter, r *http.Request) (body T, statusCode int, err error)

func ApiHandler[T any](f ApiHandlerFunc[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, statusCode, err := f(w, r)
		w.Header().Set("Content-Type", "application/json")
		if statusCode >= 200 && statusCode < 300 && err == nil {
			b, err := json.Marshal(body)
			if err != nil {
				WriteErrorResponse(w, http.StatusInternalServerError, err)
				return
			}
			if _, err = w.Write(b); err != nil {
				WriteErrorResponse(w, http.StatusInternalServerError, err)
				return
			}
		} else {
			WriteErrorResponse(w, statusCode, err)
			return
		}
	}
}
