package main

import (
	"time"
	"github.com/gorilla/websocket"
)

//cliantはチャットを行っている一人のユーザ
type client struct {
	//socketはこのクライアントのwebsocket
	socket *websocket.Conn
	//sendはメッセージが送られるチャネル
	send chan *message
	// roomはこのクライアントが参加しているチャットルーム
	room *room
	//userDataはユーザーに関する情報を保持します
	userData map[string]interface{}
}

func (c *client) read(){
	defer c.socket.Close()
	for{
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil{
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			c.room.forward <- msg
		}else{
			break
		}
	}

}

func (c *client) write(){
	defer c.socket.Close()
	for msg := range c.send{
		if err := c.socket.WriteJSON(msg); err != nil{
			return
		}
	}
}
