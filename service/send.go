package service

import (
	"Alpha-S/config"
	"Alpha-S/module/botResp"
	"Alpha-S/module/message"
	"Alpha-S/utils"
	"bytes"
	"encoding/json"
	"fmt"
)

func SendPrivateMessage(envConfig *config.Config) {
	var msg message.PrivateMessage
	msg.Message = "hello world，早上好"
	msg.UserId = 744167486
	msg.AutoEscape = true
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("SendPrivateMessage json.Marshal error:%s", err.Error())
	}

	resp, err := utils.SendBotRequest(envConfig.FrontServer.Host, "/send_private_msg", bytes.NewReader(b))
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
