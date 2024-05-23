package main

import (
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

func main() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("HTTP server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clientsMutex.Lock()
	clients[ws] = true
	clientsMutex.Unlock()

	// to hadnle the cleint disocnn
	defer func() {
		clientsMutex.Lock()
		delete(clients, ws)
		clientsMutex.Unlock()
		log.Printf("Client %s disconnected", ws.RemoteAddr())
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			// An error occurred, likely because the client disconnected
			log.Printf("Client %s disconnected: %v", ws.RemoteAddr(), err)
			break
		}
		log.Printf("recv: %s", msg)

		// Print the list of clients
		clientsMutex.Lock()
		log.Println("Current clients:")
		for client := range clients {
			log.Printf("Client address: %s", client.RemoteAddr())
		}
		clientsMutex.Unlock()

		// send message to all de peeps in teh thing
		clientsMutex.Lock()
		for client := range clients {
			if client != ws {
				err := client.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Printf("error %v:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
		clientsMutex.Unlock()

		//send mesg back to le sender
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error sending message to sender: %v", err)
		}
	}
}
