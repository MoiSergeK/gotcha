package app

import (
	socketio "github.com/googollee/go-socket.io"
	"./libs"
	// "fmt"
	// "github.com/kataras/iris"
	// "github.com/kataras/iris/websocket"
)


func InitSocketIO(io *socketio.Server){
	io.On("connection", func(socket socketio.Socket){
		libs.Println("SOCKET::connection")

		socket.On("event", func(msg string) {
			libs.Println("SOCKET::event -> " + msg)
		})

		socket.On("disconnection", func() {
			libs.Println("on disconnect")
		})
	})
}

// func SetupWebSocket(app *iris.Application) {
//     // create our echo websocket server
//     ws := websocket.New(websocket.Config{
//         ReadBufferSize:  1024,
//         WriteBufferSize: 1024,
//     })
//     ws.OnConnection(handleConnection)

//     // register the server on an endpoint.
//     // see the inline javascript code in the websockets.html,
//     // this endpoint is used to connect to the server.
//     app.Get("/echo", ws.Handler())
//     // serve the javascript built'n client-side library,
//     // see websockets.html script tags, this path is used.
//     app.Any("/iris-ws.js", websocket.ClientHandler())
// }

// func handleConnection(c websocket.Connection) {
// 	// Read events from browser
//     c.On("chat", func(msg string) {
//         // Print the message to the console, c.Context() is the iris's http context.
//         fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
//         // Write message back to the client message owner with:
//         // c.Emit("chat", msg)
//         // Write message to all except this client with:
//         c.To(websocket.Broadcast).Emit("chat", msg)
//     })
// }