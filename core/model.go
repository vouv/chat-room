package core

import "time"

// 用户在聊天室中的唯一ID
type UserID = string

const (
	EventTypeMsg    = "event-msg"    // 用户发言
	EventTypeSystem = "event-system" // 系统信息推送 如房间人数
	EventTypeJoin   = "event-join"   // 用户加入
	EventTypeTyping = "event-typing" // 用户正在输入
	EventTypeLeave  = "event-leave"  // 用户离开
	EventTypeImage  = "event-image"  // todo 消息图片
)

// 聊天室事件定义
type Event struct {
	Type      string `json:"type"`      // 事件类型
	User      string `json:"user"`      // 用户名
	Timestamp int64  `json:"timestamp"` // 时间戳
	Text      string `json:"text"`      // 事件内容
	UserCount int    `json:"userCount"` // 房间用户数
}

func NewEvent(typ string, user, msg string) Event {
	return Event{
		Type:      typ,
		User:      user,
		Timestamp: time.Now().UnixNano() / 1e6,
		Text:      msg,
	}
}

// 用户订阅
type Subscription struct {
	Id       string       // 用户在聊天室中的ID
	Username string       // 用户名
	Pipe     <-chan Event // 事件接收通道 用户从这个通道接收消息
	EmitCHn  chan Event   // 用户消息推送通道
	LeaveChn chan UserID  // 用户离开事件推送
}

func (s *Subscription) Leave() {
	s.LeaveChn <- s.Id // 将用户从聊天室列表中移除
}

func (s *Subscription) Say(message string) {
	s.EmitCHn <- NewEvent(EventTypeMsg, s.Username, message)
}
