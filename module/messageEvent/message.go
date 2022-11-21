package messageEvent

import "Alpha-S/module"

type MessageEvent struct {
	module.BaseEvent
	SubType    string `json:"sub_type"`
	MessageId  string `json:"message_id"`
	UserId     string `json:"user_id"`
	RawMessage string `json:"raw_message"`
}

type PrivateMessageEvent struct {
	MessageEvent
	RawMessage string        `json:"raw_message"`
	Font       int32         `json:"font"`
	Message    any           `json:"messag"` // todo
	Sender     PrivateSender `json:"sender"` // 根据群聊私聊mapping下面的struct
}

type GroupMessageEvent struct {
	MessageEvent
	SubType    string      `json:"sub_type"`
	SelfId     int64       `json:"self_id"`
	Anonymous  Anonymous   `json:"anonymous"`
	Message    any         `json:"messag"` // todo
	RawMessage string      `json:"raw_message"`
	Font       int32       `json:"font"`
	Sender     GroupSender `json:"sender"` // 根据群聊私聊mapping下面的struct
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

type Anonymous struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}
