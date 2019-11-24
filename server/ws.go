package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var Websocket = &ws{
	upgrader:  &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	},
}

type ws struct {
	upgrader *websocket.Upgrader
}

func (s *ws) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}

		// 加入房间
		evs := Room.GetArchive()
		Room.MsgJoin(name)
		control := Room.Join(name)
		defer control.Leave()

		// 先把历史消息推送出去
		for _, event := range evs {
			if conn.WriteJSON(&event) != nil {
				// 用户断开连接
				return
			}
		}

		// 开启通道监听用户事件然后发送给聊天室
		newMessages := make(chan string)
		go func() {
			var res = struct {
				Msg string `json:"msg"`
			}{}
			for {
				err := conn.ReadJSON(&res)
				if err != nil {
					// 用户断开连接
					close(newMessages)
					return
				}
				newMessages <- res.Msg
			}
		}()

		// 接收消息，在这里阻塞请求，循环退出就表示用户已经断开
		for {
			select {
			case event := <-control.Pipe:
				if conn.WriteJSON(&event) != nil {
					// 用户断开连接
					return
				}
			case msg, ok := <-newMessages:
				// 断开连接
				if !ok {
					return
				}
				control.Say(msg)
			}
		}
	}
}
