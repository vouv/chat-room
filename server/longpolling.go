package server

import (
	"chatroom/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var LongPolling = &longPolling{}

type longPolling struct{}

func (longPolling) Msg() gin.HandlerFunc {
	type req struct {
		Name string `json:"name"`
		Msg  string `json:"msg"`
	}
	return func(c *gin.Context) {
		form := req{}
		if err := c.BindJSON(&form); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		Room.MsgSay(form.Name, form.Msg)
		c.JSON(http.StatusOK, struct {
			Status int `json:"status"`
		}{200})
	}
}

// 轮询获取指定时间戳之后的聊天记录
func (longPolling) Archive() gin.HandlerFunc {
	return func(c *gin.Context) {
		lastReceived, _ := strconv.ParseInt(c.Query("ts"), 10, 64)

		var events []core.Event
		// filter archive
		for _, event := range Room.GetArchive() {
			if event.Timestamp > lastReceived {
				events = append(events, event)
			}
		}
		c.JSON(http.StatusOK, events)
	}
}

func (longPolling) Leave() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("name")
		Room.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
