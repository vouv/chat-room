package model

import (
	"time"
)

const (
	//msg
	EventTypeMsg = "event-msg"
	EventTypeSystem = "event-system"
	EventTypeJoin = "event-join"
	EventTypeTyping = "event-typing"
	EventTypeLeave = "event-leave"
	EventTypeImage = "event-image"
)

// 聊天室事件定义
type Event struct {
	// 事件类型
	Type string `json:"type"`
	// 用户名
	User string `json:"user"`
	// 时间戳
	Timestamp int64 `json:"timestamp"`
	// 事件内容
	Text string `json:"text"`

	UserCount int `json:"userCount"`
}

func NewEvent(typ string, user, msg string) Event {
	return Event{
		Type: typ,
		User: user,
		Timestamp: time.Now().UnixNano() / 1e6,
		Text: msg,
	}
}

// 用户订阅
type Subscription struct {
	Id string

	Username string

	// 事件接收通道
	// 用户从这个通道接收消息
	Pipe <-chan Event

	EmitCHn chan Event

	LeaveChn chan string
}

func (s *Subscription) Leave() {
	s.LeaveChn <- s.Id // 将用户从聊天室列表中移除
}

func (s *Subscription) Say(message string) {
	s.EmitCHn <- NewEvent(EventTypeMsg, s.Username, message)
}
