package model

type MessageType string

type MessageData struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}
