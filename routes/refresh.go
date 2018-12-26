package routes

import (
	"chat-room/routes/chatroom"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Refresh(server *gin.Engine)  {
	r := server.Group("/refresh")

	// Refresh
	r.GET("", func(c *gin.Context) {
		user := c.Query("user")
		droom.MsgJoin(user)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room?user=" + user)
	})


	r.GET("/room", func(c *gin.Context) {
		user := c.Query("user")

		// 只需要获取记录
		events := droom.GetArchive()

		data := struct {
			User string
			Events []chatroom.Event
		}{user,events}
		c.HTML(http.StatusOK,"refresh.html",data)
	})

	r.POST("/room", func(c *gin.Context) {
		user := c.PostForm("user")
		message := c.PostForm("message")
		droom.MsgSay(user, message)
		c.Redirect(http.StatusMovedPermanently,"/refresh/room")
	})


	r.GET("/room/leave", func(c *gin.Context) {
		user := c.Query("user")
		droom.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

}

