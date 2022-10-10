package ws

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	newline  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"wssio.dl.com"},
	}
)

func Ws(hub *hub, rw http.ResponseWriter, r *http.Request) {
	if websocket.IsWebSocketUpgrade(r) {
		conn, err := upgrader.Upgrade(rw, r, nil)
		if err != nil {
			http.Error(rw, "Not found", http.StatusNotFound)
			return
		}

		client := &client{hub: hub, conn: conn}

		go client.write()
		go client.read()
	} else {
		http.Error(rw, "Not found", http.StatusNotFound)
	}
}

func (c *client) read() {
	defer func() {
		c.conn.Close()
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("error: %v", err)
			}
			break
		}
		data = bytes.TrimSpace(bytes.Replace(data, newline, space, -1))
		logrus.Info(data)
	}
}

func (c *client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
