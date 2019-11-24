package server

import (
	"chatroom/core"
	"github.com/gin-gonic/gin"
)

var Room = core.NewRoom()

func NewServer() *gin.Engine {
	s := gin.Default()
	// static files
	s.Static("/static", "./static")
	s.StaticFile("/", "web/index.html")
	s.StaticFile("/refresh", "./web/refresh.html")
	s.StaticFile("/polling", "./web/polling.html")
	s.StaticFile("/ws", "./web/ws.html")

	{
		// refresh
		s.GET("/refresh/archive", Refresh.Archive())
		s.POST("/refresh/msg", Refresh.Msg())
		s.GET("/refresh/leave", Refresh.Leave())
	}

	{
		// polling
		s.GET("/polling/archive", LongPolling.Archive())
		s.POST("/polling/msg", LongPolling.Msg())
		s.GET("/polling/leave", LongPolling.Leave())

	}

	{
		// websocket
		s.GET("/ws/socket", Websocket.Handle())
	}

	return s
}
