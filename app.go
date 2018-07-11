package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type UsernameChange struct {
	Old string `json:"oldName"`
	New string `json:"newName"`
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		so.Join("chatroom")

		so.On("newUserJoined", func(username string) {
			// var welcomeMessage = " has joined"
			var serverMessage = username + " has joined"
			so.BroadcastTo("chatroom", "utilMessage", serverMessage)
		})

		so.On("newMessage", func(msg string) {
			var message Message
			json.Unmarshal([]byte(msg), &message)
			recordDb()
			so.BroadcastTo("chatroom", "newServerMessage", msg)
		})

		so.On("usernameChange", func(msg string) {
			var userMessage UsernameChange
			json.Unmarshal([]byte(msg), &userMessage)
			so.BroadcastTo("chatroom", "usernameChangeMessage", msg)
		})

		// so.On("newMessage", func(msg string) {
		// 	log.Println(msg)
		// 	recordDb()
		// 	log.Println("Broadcasting")
		// 	so.BroadcastTo("chatroom", "newServerMessage", msg)
		// })

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
