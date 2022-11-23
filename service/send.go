package service

import (
	"Alpha-S/module/botResp"
	"Alpha-S/module/message"
	"Alpha-S/utils"
	"bytes"
	"encoding/json"
	"fmt"
)

func SendPrivateMessage(qq int64, m string) {
	var msg message.PrivateMessage
	msg.Message = m
	msg.UserId = qq
	msg.AutoEscape = false
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("SendPrivateMessage json.Marshal error:%s", err.Error())
	}

	resp, err := utils.SendBotRequest(utils.CONFIG.FrontServer.Host, "/send_private_msg", bytes.NewReader(b))
	if err != nil {
		fmt.Printf("SendBotRequest error:%s", err.Error())
	}

	var r botResp.BotResp
	err = json.Unmarshal(resp, &r)
	if err != nil {
		fmt.Printf("SendPrivateMessage json.Unmarshal error:%s", err.Error())
	}
	println("sss")
}
