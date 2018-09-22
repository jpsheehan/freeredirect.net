package main

import (
	"strings"
	"time"
)

func stripPort(s string) string {
	i := strings.Index(s, ":")
	if i == -1 {
		return s
	}
	return s[:i]
}

func getTimeString() string {
	return time.Now().Format(time.UnixDate)
}
