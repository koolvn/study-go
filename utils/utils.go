package utils

import (
	"strings"
)

func GetBrowserName(userAgent string) string {
	var browserName string
	name := strings.ToLower(userAgent)
	switch {
	default:
		browserName = strings.Split(userAgent, " ")[len(strings.Split(userAgent, " "))-1]
	case strings.Contains(name, "chrome"):
		browserName = "Chrome"
	case strings.Contains(name, "firefox"):
		browserName = "Firefox"
	case strings.Contains(name, "opera"):
		browserName = "Opera"
	case strings.Contains(name, "edg"):
		browserName = "Edge"
	case strings.Contains(name, "safari"):
		browserName = "Safari"
	}
	return browserName
}
