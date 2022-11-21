package message

type PrivateMessage struct {
	UserId     int64  `json:"user_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type GroupMessage struct {
	GroupId    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
