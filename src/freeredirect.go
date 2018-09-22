package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jpsheehan/dotenv"
)

func main() {

	// setup the environment variables and database connection
	checkError(dotenv.Config())
	checkError(dbConnect())

	// setup the new server
	httpServer := new(http.Server)
	httpServer.Handler = NewServer().router
	httpServer.Addr = getFullAddress(os.Getenv("ADDR"), os.Getenv("PORT"))

	// handle interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Printf("%s: Shutting down...\n", getTimeString())
			dbDisconnect()
			os.Exit(0)
		}
	}()

	// print some information and serve
	fmt.Printf("%s: Listening on %s...\n", getTimeString(), httpServer.Addr)
	httpServer.ListenAndServe()
}
