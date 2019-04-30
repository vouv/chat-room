package longpolling

import (
	"chat-room/cmd/chatroom"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ROOM = chatroom.RoomS

func Msg() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := struct {
			Name string `json:"name"`
			Msg string `json:"msg"`
		}{}
		c.BindJSON(&req)
		ROOM.MsgSay(req.Name, req.Msg)
		c.JSON(http.StatusOK, struct {
			Status int `json:"status"`
		}{200})
	}
}

// 轮询获取指定时间戳之后的聊天记录
func Archive() gin.HandlerFunc {
	return func(c *gin.Context) {
		lastReceived, _ := strconv.ParseInt(c.Query("ts"), 10, 64)
		var events = []chatroom.Event{}
		for _, event := range ROOM.GetArchive() {
			if event.Timestamp > lastReceived {
				events = append(events, event)
			}
		}
		c.JSON(http.StatusOK, events)
	}
}

func Leave() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("name")
		ROOM.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func LongPolling(server *gin.Engine) {
	r := server.Group("/longpolling")

	// Long polling demo

	r.GET("/room", func(c *gin.Context) {
		user := c.Query("user")
		ROOM.MsgJoin(user)
		c.HTML(http.StatusOK, "longpolling.html", struct {
			User string
		}{user})
	})

	r.POST("/room/messages", func(c *gin.Context) {
		user := c.Query("user")
		message := c.PostForm("message")
		ROOM.MsgSay(user, message)
	})

	r.GET("/room/leave", func(c *gin.Context) {
		user := c.Query("user")
		ROOM.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/room/messages", func(c *gin.Context) {

		lastReceived, _ := strconv.ParseInt(c.Query("lastReceived"), 10, 64)

		var events = []chatroom.Event{}
		for _, event := range ROOM.GetArchive() {
			if event.Timestamp > lastReceived {
				events = append(events, event)
			}
		}

		c.JSON(http.StatusOK, events)
		return

	})

}
