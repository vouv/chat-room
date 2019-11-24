package main

import (
	"chatroom/server"
	"log"
)

func main() {
	s := server.NewServer()
	log.Fatal(s.Run())
}
