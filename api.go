package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var (
	host = "127.0.0.1"
	port = "8000"
	messages []string
)

func main() {
	router := httprouter.New()
	router.POST("/send/:msg", handlerSendMessage)

	server := http.Server{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	log.Printf("Starting server at %s:%s\n", host, port)
	server.ListenAndServe()
}

// A handler that receive string parameter and store it
func handlerSendMessage(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	messages = append(messages, p.ByName("msg"))
}