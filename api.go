package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	host = "127.0.0.1"
	port = "8000"
	messages []string
	broadcast = make(chan string)
	clients = make(map[*websocket.Conn]bool)
)

type MessageResponse struct {
	Messages []string `json:"messages"`
}

func main() {
	router := mux.NewRouter()

	//register route
	router.HandleFunc("/send/{msg}", handlerSendMessage).Methods("POST")
	router.HandleFunc("/fetch", handlerFetchAllMessage).Methods("GET")
	router.HandleFunc("/ws", handlerWebsocketConnection)
	go echo()

	server := http.Server{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	log.Printf("Starting server at %s:%s\n", host, port)
	_ = server.ListenAndServe()
}

// listen for new message being sent to server
func writeToChannel(message string) {
	broadcast <- message
}

// A handler that receive string parameter and store it
func handlerSendMessage(writer http.ResponseWriter, request *http.Request) {
	message := mux.Vars(request)["msg"]
	messages = append(messages, message)
	_, _ = fmt.Fprintf(writer, "OK")

	go writeToChannel(message)
}

// A handler to fetch all message
func handlerFetchAllMessage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	messageResponse := MessageResponse {
		Messages: messages,
	}

	err := json.NewEncoder(writer).Encode(messageResponse)
	if err != nil {
		log.Fatalf("error in encoding json: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// A handler to upgrade incoming HTTP connection
func handlerWebsocketConnection(writer http.ResponseWriter, request *http.Request) {
	client, err := websocket.Upgrade(writer, request, writer.Header(), 1024, 1024)
	if err != nil {
		log.Fatalf("error in upgrading websocket connection: %s", err)
	}

	clients[client] = true
}

// Send realtime message to all connected client
func echo() {
	for {
		message := <-broadcast

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Printf("error in write message: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}