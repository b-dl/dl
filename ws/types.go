package ws

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 1048576
)

type hub struct {
}

type client struct {
	hub *hub

	conn *websocket.Conn

	send chan []byte
}
