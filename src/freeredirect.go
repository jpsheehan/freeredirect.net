package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var hostMap map[string]string
var port int

func stripPort(s string) string {
	i := strings.Index(s, ":")
	if i == -1 {
		return s
	}
	return s[:i]
}

func getRedirectURL(hostname string) string {
	if val, ok := hostMap[hostname]; ok {
		return val
	}
	return "https://freeredirect.net/#" + hostname
}

func handler(w http.ResponseWriter, r *http.Request) {
	originalHost := stripPort(r.Host)
	redirectURL := getRedirectURL(originalHost)

	timeString := time.Now().Format(time.UnixDate)

	fmt.Printf("%s: %s -> %s\n", timeString, originalHost, redirectURL)

	w.Header().Add("Location", redirectURL)
	w.WriteHeader(302)
}

func loadHosts() {
	// Populate the hostMap
	hostMap = make(map[string]string)
	//hostMap["google.local"] = "https://google.com"
	//hostMap["sheehan.local"] = "https://sheehan.nz"
	//hostMap["freeredirect.local"] = "http://freeredirect.net"
	hostMap["auracreative.nz"] = "https://www.facebook.com/AuraCreativeArt/"
	hostMap["mclovin.co.nz"] = "https://www.linkedin.com/in/michael-gallagher-2a3538122/"
	hostMap["www.mclovin.co.nz"] = "https://www.linkedin.com/in/michael-gallagher-2a3538122/"
}

func main2() {
	loadHosts()

	port = 8081
	address := "0.0.0.0"

	timeString := time.Now().Format(time.UnixDate)
	fullAddress := address + ":" + strconv.Itoa(port)

	fmt.Printf("%s: Listening on %s...\n", timeString, fullAddress)

	http.HandleFunc("/", handler)
	http.ListenAndServe(fullAddress, nil)
}
