package routes

import (
	"chat-room/routes/chatroom"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LongPolling(server *gin.Engine)  {
	r := server.Group("/longpolling")

	// Long polling demo

	r.GET("/room", func(c *gin.Context) {
		user := c.Query("user")
		droom.MsgJoin(user)
		c.HTML(http.StatusOK,"longpolling.html", struct {
			User string
		}{user})
	})


	r.POST("/room/messages", func(c *gin.Context) {
		user := c.Query("user")
		message := c.PostForm("message")
		droom.MsgSay(user, message)
	})

	r.GET("/room/leave", func(c *gin.Context) {
		user := c.Query("user")
		droom.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/room/messages", func(c *gin.Context) {

		lastReceived,_ := strconv.ParseInt(c.Query("lastReceived"),10,64)

		var events = []chatroom.Event{}
		for _, event := range droom.GetArchive() {
			if event.Timestamp > lastReceived {
				events = append(events, event)
			}
		}

		c.JSON(http.StatusOK,events)
		return

	})


}





