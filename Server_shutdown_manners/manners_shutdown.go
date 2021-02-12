package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

// Empty struct
type handler struct{}

// Struct's method to response
func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Pooper Poop"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

// Start
func main() {
	handler := newHandler()                  // create instance to call server for response
	ch := make(chan os.Signal)               // create channel
	signal.Notify(ch, os.Interrupt, os.Kill) // send shutdown from OS to ch
	go listenForShutdown(ch)                 // wait for signal from channel

	manners.ListenAndServe(":8080", handler) // start server
}

func newHandler() *handler {
	return &handler{}
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch // block channel while waiting for signal
	// continue execution when channel'll close
	fmt.Println("Server will stop now")
	manners.Close()
}
