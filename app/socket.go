package app

import (
	socketio "github.com/googollee/go-socket.io"
	"./libs"
)


func InitSocketIO(io *socketio.Server){
	io.On("connection", func(socket socketio.Socket){
		libs.Println("SOCKET::connection")

		socket.On("event", func(msg string) {
			libs.Println("SOCKET::event -> " + msg)
		})
	})
}