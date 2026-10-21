package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func getTitleAndIDFromSlug(slug string) (string, int64) {
	lastHyphenIndex := strings.LastIndex(slug, "-")
	if lastHyphenIndex == -1 {
		id, err := strconv.ParseInt(slug, 10, 64)
		if err != nil {
			return "", 0
		}

		return "", id
	}

	id, err := strconv.ParseInt(slug[lastHyphenIndex+1:], 10, 64)
	if err != nil {
		return "", 0
	}

	title := slug[:lastHyphenIndex]

	return title, id
}

func makeSlug(input string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}

	processedString := reg.ReplaceAllString(input, " ")
	processedString = strings.TrimSpace(processedString)

	slug := strings.ReplaceAll(processedString, " ", "-")
	slug = strings.ToLower(slug)

	return slug, nil
}

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
	if wantsJSON {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": message,
		})
		return
	}

	http.Error(w, message, status)
}
