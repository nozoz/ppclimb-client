package main

import (
	"os"
	"os/signal"
	"ppclimb-client/internal/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := websocket.NewClient()
	go c.Run()

	<-interrupt
}
