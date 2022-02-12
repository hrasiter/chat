package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	socket *websocket.Conn
	room   *Room
	send   chan []byte
}

func (c *Client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()

		if err != nil {
			fmt.Println("func (c *Client) read(): Error reading socket message!!!")
			return
		}

		c.room.forward <- msg
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			fmt.Println("func (c *Client) write(): Error in writing into channel")
		}
	}
}
