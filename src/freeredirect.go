package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jpsheehan/dotenv"
)

var port int

func stripPort(s string) string {
	i := strings.Index(s, ":")
	if i == -1 {
		return s
	}
	return s[:i]
}

func handler(w http.ResponseWriter, r *http.Request) {
	originalHost := stripPort(r.Host)
	redirectURL := dbGetRedirectURL(originalHost)

	timeString := time.Now().Format(time.UnixDate)

	fmt.Printf("%s: %s -> %s\n", timeString, originalHost, redirectURL)

	w.Header().Add("Location", redirectURL)
	w.WriteHeader(302)
}

func main() {

	checkError(dotenv.Config())
	checkError(dbConnect())

	port, err := strconv.Atoi(os.Getenv("PORT"))
	address := os.Getenv("ADDR")

	checkError(err)

	timeString := time.Now().Format(time.UnixDate)
	fullAddress := address + ":" + strconv.Itoa(port)

	fmt.Printf("%s: Listening on %s...\n", timeString, fullAddress)

	http.HandleFunc("/", handler)
	http.ListenAndServe(fullAddress, nil)
}
