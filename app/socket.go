package app

import (
	socketio "github.com/googollee/go-socket.io"
)


func InitSocketIO(io *socketio.Server){
	io.On("connection", func(socket socketio.Socket){
		
	})
}