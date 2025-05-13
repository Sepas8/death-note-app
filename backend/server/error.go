package server

import (
	"encoding/json"
	"net/http"

	"import github.com/Sepas8/death-note-app/backend/api"
)

var statusMap = map[int]string{
	400: "Bad Request",
	404: "Not Found",
	405: "Method Not Allowed",
	500: "Internal Server Error",
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",
}

func (s *Server) HandleError(w http.ResponseWriter, statusCode int, path string, cause error) {
	var errorResponse api.ErrorResponse
	errorResponse.Status = statusCode
	errorResponse.Message = cause.Error()
	errorResponse.Description = statusMap[statusCode]
	response, err := json.Marshal(errorResponse)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, path, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
	s.logger.Error(statusCode, path, cause)
}
