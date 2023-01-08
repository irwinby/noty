package status

import (
	"encoding/json"
	"net/http"
)

type (
	ErrorType uint

	ErrorResponseDetail struct {
		Message string `json:"message"`
	}

	ErrorResponse struct {
		Message string                `json:"message"`
		Code    int                   `json:"code"`
		Details []ErrorResponseDetail `json:"details"`
	}
)

const (
	InternalServerError = iota + 1
	BadRequestError
	NotFoundError
)

func (t ErrorType) Code() int {
	switch t {
	case InternalServerError:
		return http.StatusInternalServerError
	case BadRequestError:
		return http.StatusBadRequest
	case NotFoundError:
		return http.StatusNotFound
	default:
		return 0
	}
}

func (t ErrorType) Message() string {
	switch t {
	case InternalServerError:
		return "Internal Server Error"
	case BadRequestError:
		return "Bad Request Error"
	case NotFoundError:
		return "Not Found Error"
	default:
		return ""
	}
}

func Error(w http.ResponseWriter, err ErrorType, details ...string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(err.Code())

	response := ErrorResponse{
		Message: err.Message(),
		Code:    err.Code(),
	}

	for _, detail := range details {
		response.Details = append(response.Details, ErrorResponseDetail{
			Message: detail,
		})
	}

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(&response)
}

func OK(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(&body)
}

func Empty(w http.ResponseWriter, code int) {
	OK(w, "[]", code)
}
