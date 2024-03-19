package model

type MessageType string

type MessageData struct {
	Type    string `json:"type"`
	Author  string `json:"author"`
	Payload string `json:"payload"`
}
