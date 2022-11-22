package service

import (
	"Alpha-S/module"
	"Alpha-S/module/messageEvent"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func GetEvent(c *gin.Context) {
	var event module.BaseEvent
	//err := c.BindJSON(&event)
	requestBody, err := io.ReadAll(c.Request.Body)
	s := string(requestBody)
	println(s)
	if err != nil {
		fmt.Printf("GetEvent BindJSON error:%s \n", err.Error())
	}
	err = json.Unmarshal(requestBody, &event)
	switch event.PostType {
	case "message":
		var me messageEvent.MessageEvent
		err = json.Unmarshal(requestBody, &me)
		if err != nil {
			fmt.Printf("MessageEvent BindJSON error:%s \n", err.Error())
		}
		if me.SubType == "friend" {
			var pme messageEvent.PrivateMessageEvent
			err = json.Unmarshal(requestBody, &pme)
			handlePMessage(&pme)
		} else if me.SubType == "group" {
			var gme messageEvent.GroupMessageEvent
			err = json.Unmarshal(requestBody, &gme)
		}
		if err != nil {
			fmt.Printf("P/G MessageEvent BindJSON error:%s \n", err.Error())
		}
	}
}

func SendTest(c *gin.Context) {
	//SendPrivateMessage(utils.CONFIG)
}

func handlePMessage(pme *messageEvent.PrivateMessageEvent) {
	info := PrivateLogic(pme.Message)
	SendPrivateMessage(pme.UserId, info)
}
