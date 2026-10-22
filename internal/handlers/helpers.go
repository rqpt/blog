package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func acceptsJSON(r *http.Request) bool {
	return strings.Contains(
		r.Header.Get("Accept"),
		"application/json",
	)
}

func respondError(
	w http.ResponseWriter,
	message string,
	status int,
	wantsJSON bool,
) {
	if !wantsJSON {
		http.Error(w, message, status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
