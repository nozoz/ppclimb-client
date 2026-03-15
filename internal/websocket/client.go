package websocket

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connected bool
	url       url.URL
	conn      *websocket.Conn
}

func NewClient() *Client {
	return &Client{
		url: url.URL{Scheme: "ws", Host: "127.0.0.1:24050", Path: "/websocket/v2"},
	}
}

func (c *Client) Connect() error {
	log.Printf("connecting to %s", c.url.String())
	conn, _, err := websocket.DefaultDialer.Dial(c.url.String(), nil)
	if err != nil {
		log.Printf("dial error: %s", err)
		return err
	}

	c.conn = conn
	c.Connected = true
	log.Printf("connected!")
	return nil
}

func (c *Client) Run() {
	for {
		if err := c.Connect(); err != nil {
			log.Println("Attempting re-connection in 3 seconds...")
			time.Sleep(3 * time.Second)
			continue
		}
	}
}
