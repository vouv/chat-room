package routes

import (
	"chat-room/routes/chatroom"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var droom = chatroom.RoomS

func WebSocket(server *gin.Engine)  {

	r := server.Group("/websocket")

	r.GET("/room", func(c *gin.Context) {
		user := c.Query("user")
		c.HTML(http.StatusOK,"websocket.html", struct {
			User string
		}{user})
	})


	r.GET("/room/socket", func(c *gin.Context) {
		user := c.Query("user")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}

		// 加入房间
		evs := droom.GetArchive()
		droom.MsgJoin(user)
		ctl := droom.Join(user)
		defer ctl.Leave()

		//先把历史消息推送出去
		for _, event := range evs {
			if conn.WriteJSON(&event) != nil {
				// They disconnected
				return
			}
		}

		newMessages := make(chan string)
		go func() {
			var res = struct {
				Msg string `json:"msg"`
			}{}
			for {
				err := conn.ReadJSON(&res)
				if err != nil {
					close(newMessages)
					return
				}
				newMessages <- res.Msg
			}
		}()

		// 接收消息
		for {
			select {
			case event := <- ctl.Pipe:
				if conn.WriteJSON(&event) != nil {
					// 断开
					return
				}
			case msg, ok := <- newMessages:
				// 断开连接
				if !ok {
					return
				}
				ctl.Say(msg)
			}
		}
	})
}




