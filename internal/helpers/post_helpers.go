package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ParseID(slug string) (int64, error) {
	lastHyphenIndex := strings.LastIndex(slug, "-")

	idPart := slug
	if lastHyphenIndex != -1 {
		idPart = slug[lastHyphenIndex+1:]
	}

	id, err := strconv.ParseInt(idPart, 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid slug ID")
	}

	return id, nil
}

func AcceptsJSON(r *http.Request) bool {
	return strings.Contains(
		r.Header.Get("Accept"),
		"application/json",
	)
}

func RespondError(
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

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
