/**
* Created by cks
* Date: 2020-12-02
* Time: 15:39
*/
package connect

import (
	"SLU-System/proto"
	"github.com/gorilla/websocket"
)

//in fact, Channel it's a user Connect session
type Channel struct {
	Room      *Room
	Next      *Channel
	Prev      *Channel
	broadcast chan *proto.Msg
	userId    int
	conn      *websocket.Conn
}

func NewChannel(size int) (c *Channel) {
	c = new(Channel)
	c.broadcast = make(chan *proto.Msg, size)
	c.Next = nil
	c.Prev = nil
	return
}

func (ch *Channel) Push(msg *proto.Msg) (err error) {
	select {
	case ch.broadcast <- msg:
	default:
	}
	return
}
