package messageEvent

import "Alpha-S/module"

// MessageEvent todo sub message struct
type MessageEvent struct {
	module.Event
	//MessageType string `json:"message_type"`
	SubType    string `json:"sub_type"`
	MessageId  string `json:"message_id"`
	UserId     string `json:"user_id"`
	RawMessage string `json:"raw_message"`
	Font       int32  `json:"font"`
	Sender     string `json:"sender"` // 根据群聊私聊mapping下面的struct
}

type GroupSender struct {
	UserId   string `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Card     string `json:"card"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}

type PrivateSender struct {
	UserId   string `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}
