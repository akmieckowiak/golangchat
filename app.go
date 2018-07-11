package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

type Message struct {
	username string `json:"username"`
	content  string `json:"content"`
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	var welcomeMessage = " has joined"

	server.On("connection", func(so socketio.Socket) {
		so.Join("chatroom")

		so.On("newUserJoined", func(username string) {
			log.Println(username)
			var serverMessage = username + welcomeMessage
			so.BroadcastTo("chatroom", "utilMessage", serverMessage)
		})

		so.On("newMessage", func(msg string) {
			log.Println(msg)
			recordDb()
			log.Println("Broadcasting")
			so.BroadcastTo("chatroom", "newServerMessage", msg)
		})

		// so.On("chat message", func(msg string) {
		// 	log.Println("emit:", so.Emit("chat message", msg))
		// 	so.BroadcastTo("chatroom", "chat message", msg)
		// })

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func recordDb() {
	log.Println("Saving to db")
}
