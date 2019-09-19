package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var (
	host = "127.0.0.1"
	port = "8000"
)

func main() {
	router := httprouter.New()

	server := http.Server{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}
	server.ListenAndServe()
	return
}
