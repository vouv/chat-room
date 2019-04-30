package chatroom

import (
	"time"
)

const (
	//msg
	EVENT_TYPE_MSG = iota
	EVENT_TYPE_JOIN
	EVENT_TYPE_TYPING
	EVENT_TYPE_LEAVE
	EVENT_TYPE_IMAGE
)

// 聊天室事件定义
type Event struct {
	// 事件类型
	Type int 		`json:"type"`
	// 用户名
	User string 	`json:"user"`
	// 时间戳
	Timestamp int64 `json:"timestamp"`
	// 事件内容
	Text string 	`json:"text"`
}

func newEvent(typ int, user, msg string) Event {
	return Event{typ, user, time.Now().UnixNano()/1e6, msg}
}

// 用户订阅
type Subscription struct {
	id int64

	username string

	// 事件接收通道
	// 用户从这个通道接收消息
	Pipe <-chan Event

	emit chan Event

	leave chan int64
}

func (s *Subscription) Leave() {
	s.leave <- s.id // 将用户从聊天室列表中移除
}

func (s *Subscription) Say(message string) {
	s.emit <- newEvent(EVENT_TYPE_MSG, s.username, message)
}
