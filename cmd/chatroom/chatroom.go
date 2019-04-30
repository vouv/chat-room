package chatroom

import (
	"container/list"
)

// 保存历史消息的条数
const archiveSize = 20
const chanSize = 10
const MSG_JOIN = "[加入房间]"
const MSG_LEAVE = "[离开房间]"
const MSG_TYPING = "[正在输入]"

var RoomS *Room

// 聊天室
type Room struct {
	// 订阅者列表
	users     map[int64]chan Event
	userCount int64
	// 自增id作为用户索引
	idx int64

	// 聊天室的消息推送入口
	publishChn chan Event

	// 历史记录
	archive *list.List

	archiveChan chan chan []Event

	// 接收订阅事件的通道
	// 用户加入聊天室后要把历史事件推送给用户
	joinChn chan chan Subscription

	// 用户取消订阅通道
	// 把通道中的历史事件释放
	// 并把用户从聊天室用户列表中删除
	leaveChn chan int64
}

func NewRoom() *Room {
	r := &Room{
		users:     map[int64]chan Event{},
		userCount: 0,
		idx:       0,

		publishChn:  make(chan Event, chanSize),
		archiveChan: make(chan chan []Event, chanSize),
		archive:     list.New(),

		joinChn:  make(chan chan Subscription, chanSize),
		leaveChn: make(chan int64, chanSize),
	}

	go r.Serve()

	return r
}

// 用来向聊天室发送用户消息
// 这些接口供非websocket连接方式调用
func (r *Room) MsgJoin(user string) {
	r.publishChn <- newEvent(EVENT_TYPE_JOIN, user, MSG_JOIN)
}

func (r *Room) MsgSay(user, message string) {
	r.publishChn <- newEvent(EVENT_TYPE_MSG, user, message)
}

func (r *Room) MsgLeave(user string) {
	r.publishChn <- newEvent(EVENT_TYPE_MSG, user, MSG_LEAVE)
}

func (r *Room) Remove(id int64) {
	r.leaveChn <- id // 将用户从聊天室列表中移除
}

// 用户订阅聊天室入口函数
// 返回用户订阅的对象，用户根据对象中的属性读取历史消息和即时消息
func (r *Room) Join(username string) Subscription {
	resp := make(chan Subscription)
	r.joinChn <- resp
	s := <-resp
	s.username = username
	return s
}

func (r *Room) GetArchive() []Event {
	ch := make(chan []Event)
	r.archiveChan <- ch
	return <-ch
}

// 处理聊天室中的事件
func (r *Room) Serve() {
	for {
		select {
		// 用户加入房间
		case ch := <-r.joinChn:
			chE := make(chan Event, chanSize)
			r.userCount++
			r.idx++
			r.users[r.idx] = chE
			ch <- Subscription{
				id:    r.idx,
				Pipe:  chE,
				emit:  r.publishChn,
				leave: r.leaveChn,
			}
		case arch := <-r.archiveChan:
			events := []Event{}
			//历史事件
			for e := r.archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(Event))
			}
			arch <- events
		// 有新的消息
		case event := <-r.publishChn:
			// 推送给所有用户
			for _, v := range r.users {
				v <- event
			}
			// 推送消息后，限制本地只保存指定条历史消息
			if r.archive.Len() >= archiveSize {
				r.archive.Remove(r.archive.Front())
			}
			r.archive.PushBack(event)
		// 用户退出房间
		case k := <-r.leaveChn:
			if _, ok := r.users[k]; ok {
				delete(r.users, k)
				r.userCount--
			}
		}
	}
}

// 开启goroutine loop Serve
func init() {
	RoomS = NewRoom()
}
