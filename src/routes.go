package main

import (
	"fmt"
	"net/http"
	"os"
)

// handles what happens when a host that isn't freeredirect.net accesses the site
func (s *Server) handleDefaultHosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalHost := stripPort(r.Host)
		redirectURL := dbGetRedirectURL(originalHost)

		fmt.Printf("%s: %s -> %s\n", getTimeString(), originalHost, redirectURL)

		w.Header().Add("Location", redirectURL)
		w.WriteHeader(302)
	}
}

func (s *Server) handleStaticDocument(staticRoot string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := staticRoot + r.URL.Path
		if _, err := os.Stat(filename); err == nil {
			file, err := os.Open(filename)
			file.Read() // finish!!!
		}
	}
}
