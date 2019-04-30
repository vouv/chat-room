package chatroom

import (
	"fmt"
	"testing"
	"time"
)

func TestJoin(t *testing.T) {

	room := RoomS

	jack := room.Join("jack")
	go func() {
		for v := range jack.Pipe {
			fmt.Println(v)
		}
	}()

	tom := room.Join("Tom")
	go func() {
		for v := range tom.Pipe {
			fmt.Println(v)
		}
	}()
	jack.Say("hello world")
	tom.Say("nice to meet U")

	jack.Leave()
	tom.Leave()
	time.Sleep(1 * time.Second)

}
