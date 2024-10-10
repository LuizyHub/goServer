// Package handlers handlers/response.go
package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse 에러 응답을 나타냅니다.
type ErrorResponse struct {
	Message string `json:"message"`
}

// WriteErrorResponse 에러 응답을 작성하는 헬퍼 함수입니다.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
