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

func main() {

	checkError(dotenv.Config())
	checkError(dbConnect())

	port, err := strconv.Atoi(os.Getenv("PORT"))
	address := os.Getenv("ADDR")

	checkError(err)

	timeString := time.Now().Format(time.UnixDate)
	fullAddress := address + ":" + strconv.Itoa(port)

	fmt.Printf("%s: Listening on %s...\n", timeString, fullAddress)

	s := NewServer()
	httpServer := new(http.Server)
	httpServer.Handler = s.router
	httpServer.Addr = fullAddress
	httpServer.ListenAndServe()
}
