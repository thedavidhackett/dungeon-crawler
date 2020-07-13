package websocket

import (
	"go.mongodb.org/mongo-driver/mongo"
	"dungeon-crawler/models"
)

type message struct {
	data models.Dungeon
	room string
}

//Subscription ..
type Subscription struct {
	Conn *Connection
	Room string
	Db mongo.Database
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	Rooms map[string]map[*Connection]bool

	// Inbound messages from the connections.
	Broadcast chan message

	// Register requests from the connections.
	Register chan Subscription

	// Unregister requests from connections.
	Unregister chan Subscription
}

//Hub ..
var Hub = hub{
	Broadcast:  make(chan message),
	Register:   make(chan Subscription),
	Unregister: make(chan Subscription),
	Rooms:      make(map[string]map[*Connection]bool),
}

//Run ..
func (h *hub) Run() {
	for {
		select {
		case s := <-h.Register:
			connections := h.Rooms[s.Room]
			if connections == nil {
				connections = make(map[*Connection]bool)
				h.Rooms[s.Room] = connections
			}
			h.Rooms[s.Room][s.Conn] = true
		case s := <-h.Unregister:
			connections := h.Rooms[s.Room]
			if connections != nil {
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
					close(s.Conn.Send)
					if len(connections) == 0 {
						delete(h.Rooms, s.Room)
					}
				}
			}
		case m := <-h.Broadcast:
			connections := h.Rooms[m.room]
			for c := range connections {
				select {
				case c.Send <- m.data:
				default:
					close(c.Send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.Rooms, m.room)
					}
				}
			}
		}
	}
}
