package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/jpsheehan/dotenv"
)

func main() {

	checkError(dotenv.Config())
	checkError(dbConnect())

	port, err := strconv.Atoi(os.Getenv("PORT"))
	address := os.Getenv("ADDR")

	checkError(err)

	timeString := time.Now().Format(time.UnixDate)
	fullAddress := address + ":" + strconv.Itoa(port)

	fmt.Printf("%s: Listening on %s...\n", timeString, fullAddress)

	// setup the new server
	httpServer := new(http.Server)
	httpServer.Handler = NewServer().router
	httpServer.Addr = fullAddress

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Printf("%s: Shutting down...\n", getTimeString())
			dbDisconnect()
			os.Exit(0)
		}
	}()

	httpServer.ListenAndServe()
}
