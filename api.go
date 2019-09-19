package main

import (
	"encoding/json"
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

type MessageResponse struct {
	Message []string `json:"message"`
}

func main() {
	router := httprouter.New()
	router.POST("/send/:msg", handlerSendMessage)
	router.GET("/fetch", handlerFetchAllMessage)

	server := http.Server{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	log.Printf("Starting server at %s:%s\n", host, port)
	_ = server.ListenAndServe()
}

// A handler that receive string parameter and store it
func handlerSendMessage(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	messages = append(messages, p.ByName("msg"))
	_, _ = fmt.Fprintf(writer, "OK")
}

// A handler to fetch all message
func handlerFetchAllMessage(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	messageResponse := MessageResponse {
		Message: messages,
	}
	err := json.NewEncoder(writer).Encode(messageResponse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
