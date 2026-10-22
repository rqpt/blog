package slug

import (
	"regexp"
	"strconv"
	"strings"
)

var nonAlphanumeric = regexp.MustCompile("[^a-zA-Z0-9]+")

func Parse(slug string) (string, int64) {
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

func Make(input string) string {
	processed := nonAlphanumeric.ReplaceAllString(input, " ")
	processed = strings.TrimSpace(processed)
	processed = strings.ReplaceAll(processed, " ", "-")

	return strings.ToLower(processed)
}
