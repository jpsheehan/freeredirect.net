package main

import (
	"fmt"
	"strings"
	"time"
)

func getFullAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

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
