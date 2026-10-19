package handlers

import (
	"strconv"
	"strings"
)

func getIDFromSlug(slug string) int64 {
	i := strings.LastIndex(slug, "-")
	if i == -1 {
		return 0
	}

	id, err := strconv.ParseInt(slug[i+1:], 10, 64)
	if err != nil {
		return 0
	}

	return id
}
