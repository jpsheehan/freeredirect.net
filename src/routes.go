package main

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) handleDefaultHosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalHost := stripPort(r.Host)
		redirectURL := dbGetRedirectURL(originalHost)

		timeString := time.Now().Format(time.UnixDate)

		fmt.Printf("%s: %s -> %s\n", timeString, originalHost, redirectURL)

		w.Header().Add("Location", redirectURL)
		w.WriteHeader(302)
	}
}
