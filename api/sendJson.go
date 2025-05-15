package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJson(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error(
			"Failed to marshal json data",
			"erro", err,
		)
		SendJson(
			w,
			Response{Error: "Something is wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error(
			"Error to write response to client",
			"error", err,
		)
		return
	}
}
