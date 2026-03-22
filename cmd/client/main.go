package main

import (
	"os"
	"os/signal"
	"ppclimb-client/internal/tracker"
	"ppclimb-client/internal/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := websocket.NewClient()
	go c.Run()

	t := tracker.NewTracker(c.Messages)
	go t.Run()

	<-interrupt
}
