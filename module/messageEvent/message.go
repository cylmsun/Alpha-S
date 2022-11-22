package messageEvent

import "Alpha-S/module"

type MessageEvent struct {
	module.BaseEvent
	SubType    string `json:"sub_type"` // 如果是好友则是 friend, 如果是群临时会话则是 group, 如果是在群中自身发送则是 group_self, 正常群聊消息是 normal, 匿名消息是 anonymous, 系统提示 ( 如「管理员已禁止群内匿名聊天」 ) 是 notice
	MessageId  int32  `json:"message_id"`
	UserId     int64  `json:"user_id"`
	RawMessage string `json:"raw_message"`
}

type PrivateMessageEvent struct {
	MessageEvent
	RawMessage string        `json:"raw_message"`
	Font       int32         `json:"font"`
	Message    string        `json:"message"` // todo
	Sender     PrivateSender `json:"sender"`  // 根据群聊私聊mapping下面的struct
}

type GroupMessageEvent struct {
	MessageEvent
	SelfId     int64       `json:"self_id"`
	Anonymous  Anonymous   `json:"anonymous"`
	Message    any         `json:"message"` // todo
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
	UserId   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type Anonymous struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}
