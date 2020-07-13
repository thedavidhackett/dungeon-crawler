package websocket

import (
	"fmt"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/websocket"

	"dungeon-crawler/models"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

//Upgrader ..
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	Ws *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan models.Dungeon
}

// ReadPump pumps messages from the websocket connection to the hub.
func (s Subscription) ReadPump() {
	c := s.Conn
	defer func() {
		Hub.Unregister <- s
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var d models.Dungeon
		err := c.Ws.ReadJSON(&d)
		if err != nil {
			fmt.Println(err.Error())
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		collection := s.Db.Collection("dungeons")
		_, err = collection.ReplaceOne(context.TODO(), bson.M{"_id": d.ID}, d)
		if (err != nil) {
			fmt.Println(err.Error())
		}
		fmt.Println("something")
		m := message{d, s.Room}
		Hub.Broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(mt, payload)
}

// write writes a message with the given message type and payload.
func (c *Connection) writeJSON(payload models.Dungeon) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteJSON(payload)
}

// WritePump pumps messages from the hub to the websocket connection.
func (s *Subscription) WritePump() {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.writeJSON(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// // serveWs handles websocket requests from the peer.
// func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
// 	fmt.Print(roomId)
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	c := &connection{send: make(chan []byte, 256), ws: ws}
// 	s := subscription{c, roomId}
// 	h.register <- s
// 	go s.writePump()
// 	go s.readPump()
// }
