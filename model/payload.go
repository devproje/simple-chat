package model

type MessageType string

const (
	NewMessage  MessageType = "new_message"
	SetUsername MessageType = "set_username"
	LeftUser    MessageType = "left_user"
)

type MessageData struct {
	Type    MessageType `json:"type"`
	Author  string      `json:"author"`
	Payload string      `json:"payload"`
}

func (t MessageType) ToString() string {
	return string(t)
}
