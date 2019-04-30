package refresh

import (
	"chat-room/cmd/chatroom"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ROOM = chatroom.RoomS

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		ROOM.MsgJoin(user)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room?user="+user)
	}
}

func Room() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")

		// 只需要获取记录
		events := ROOM.GetArchive()

		data := struct {
			User   string
			Events []chatroom.Event
		}{user, events}
		c.HTML(http.StatusOK, "refresh.html", data)
	}
}

func Msg() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.PostForm("user")
		message := c.PostForm("message")
		ROOM.MsgSay(user, message)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room")
	}
}

func Leave() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		ROOM.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
