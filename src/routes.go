package main

import (
	"fmt"
	"net/http"
)

func (s *Server) handleDefaultHosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalHost := stripPort(r.Host)
		redirectURL := dbGetRedirectURL(originalHost)

		fmt.Printf("%s: %s -> %s\n", getTimeString(), originalHost, redirectURL)

		w.Header().Add("Location", redirectURL)
		w.WriteHeader(302)
	}
}
