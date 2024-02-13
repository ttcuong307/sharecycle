package web

import (
	"encoding/json"
	"net/http"
	"sharecycle/foundation/cerror"
)

type ErrorResponse struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}

func RespondError(w http.ResponseWriter, err error) error {
	// TODO get statuscode base on error
	if err := Respond(w, ErrorResponse{
		Error:   err.Error(),
		Details: []string{cerror.UserMessage(err)},
	}, cerror.StatusCode(err)); err != nil {
		return err
	}

	return nil
}

func Respond(w http.ResponseWriter, resp interface{}, statusCode int) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent || resp == nil {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
