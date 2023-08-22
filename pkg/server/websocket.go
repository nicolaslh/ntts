package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nicolaslh/ntts/pkg/lib/proto"
)

var updrager = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //allow all connect
	},
}

func NewWebsocket() {
	http.HandleFunc("/ws", handleConnection)
	log.Println("Websocket server started on :8901")
	err := http.ListenAndServe(":8901", nil)
	if err != nil {
		panic(err)
	}
	
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := updrager.Upgrade(w, r, nil)
	if err != nil {
		log.Panic("Error upgrading connect: ", err)
	}


	log.Println("Client connected")

	msg := make(chan []byte, 10)
	go process(conn, msg)
	go listen(conn, msg)

}

func process(conn *websocket.Conn, msgch <-chan []byte) {
	for msg := range msgch {
		var r proto.RequestMessage
		err := json.Unmarshal(msg, &r)
		if err != nil {
			log.Printf("Error json unmarshal: %v", err)
			continue
		}

		switch r.Command {
		case "clock in":
			log.Printf("Received clock in message")
		}
	}
}

func listen(conn *websocket.Conn, msgch chan<- []byte) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		msgch <- msg
		log.Printf("Received message: %d, messageType: %s", messageType, msg)
	}
	conn.Close()
	log.Println("Client disconnected")
}
