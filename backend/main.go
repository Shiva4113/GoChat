// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// func reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Host)
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	reader(ws)
// }

// func setupRoutes() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Simple Server")
// 	})
// 	http.HandleFunc("/ws", serveWs)
// }

// func main() {
// 	fmt.Println("Chat App v0.01")
// 	setupRoutes()
// 	http.ListenAndServe(":8080", nil)
// }

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

	log.Println("http server started on :8080")
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

	//handling client diconenction :-
	defer func() {
		clientsMutex.Lock()
		delete(clients, ws)
		clientsMutex.Unlock()
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		log.Printf("recv: %s", msg)

		// err = ws.WriteMessage(websocket.TextMessage, msg)
		// if err != nil {
		// 	log.Printf("error: %v", err)
		// 	break
		// }

		clientsMutex.Lock()
		for client := range clients {
			if client != ws {
				err := ws.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Printf("error %v:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
		clientsMutex.Unlock()
	}
}
