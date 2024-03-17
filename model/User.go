package model

import "github.com/gorilla/websocket"

type User struct {
	Sock *websocket.Conn
	Name string
}
