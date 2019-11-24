package server

import (
	"chatroom/core"
	"github.com/gin-gonic/gin"
	"net/http"
)


var Refresh = &refresh{}
type refresh struct {}

func (refresh) Index() gin.HandlerFunc {

	return func(c *gin.Context) {
		user := c.Query("user")
		Room.MsgJoin(user)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room?user="+user)
	}

}

func (refresh) Archive() gin.HandlerFunc {
	type archive struct {
		User   string
		Events []core.Event
	}
	return func(c *gin.Context) {
		user := c.Query("user")

		c.HTML(http.StatusOK, "refresh.html", archive{
			User: user,
			Events: Room.GetArchive(),
		})
	}
}

func (refresh) Msg() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.PostForm("user")
		message := c.PostForm("message")
		Room.MsgSay(user, message)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room")
	}
}

func (refresh) Leave() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		Room.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
