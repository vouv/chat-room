package main

import (
	"chat-room/cmd/longpolling"
	"chat-room/cmd/refresh"
	"chat-room/cmd/ws"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()

	// index
	app.StaticFile("/", "./web/index.html")

	// refresh
	app.StaticFile("/refresh", "./web/refresh.html")
	app.GET("/refresh/archive", refresh.Room())
	app.POST("/refresh/msg", refresh.Msg())
	app.GET("/refresh/leave", refresh.Leave())

	// polling
	app.StaticFile("/polling", "./web/polling.html")
	app.GET("/polling/archive", longpolling.Archive())
	app.POST("/polling/msg", longpolling.Msg())
	app.GET("/polling/leave", longpolling.Leave())

	// websocket
	app.StaticFile("/ws", "./web/ws.html")
	app.GET("/ws/socket", ws.Socket())

	// static files
	app.Static("/static", "./static")

	log.Fatal(app.Run())
}
