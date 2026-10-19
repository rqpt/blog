package handlers

import (
	"strconv"
	"strings"
)

func getTitleAndIDFromSlug(slug string) (string, int64) {
	lastHyphenIndex := strings.LastIndex(slug, "-")
	if lastHyphenIndex == -1 {
		return "", 0
	}

	id, err := strconv.ParseInt(slug[lastHyphenIndex+1:], 10, 64)
	if err != nil {
		return "", 0
	}

	title := slug[:lastHyphenIndex]

	return title, id
}

