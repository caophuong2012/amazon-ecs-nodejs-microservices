package response

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

// ErrorResponse
type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// RespondJSON writes JSON as http response
func RespondJSON(w http.ResponseWriter, httpStatusCode int, object interface{}, headers map[string]string) {
	bytes, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key, value := range headers {
		w.Header().Set(key, value)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatusCode)
	w.Write(bytes)
}

func ErrorJSON(ctx context.Context, w http.ResponseWriter, httpStatusCode int, err string, description string) {

	log.Error().Msg(fmt.Sprintf("Error: %v, Response Code: %v, Description: %v", err, httpStatusCode, description))

	RespondJSON(w, httpStatusCode, ErrorResponse{
		Error:            err,
		ErrorDescription: description,
	}, nil)
}
