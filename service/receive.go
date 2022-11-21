package service

import (
	"Alpha-S/module"
	"Alpha-S/module/messageEvent"
	"Alpha-S/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	var event module.BaseEvent
	err := c.BindJSON(&event)
	if err != nil {
		fmt.Printf("GetEvent BindJSON error:%s \n", err.Error())
	}
	switch event.PostType {
	case "message":
		var me messageEvent.MessageEvent
		err = c.BindJSON(&me)
		if err != nil {
			fmt.Printf("MessageEvent BindJSON error:%s \n", err.Error())
		}
		if me.SubType == "public" {
			var pme messageEvent.PrivateMessageEvent
			err = c.BindJSON(&pme)
			handlePMessage(&pme)
		} else if me.SubType == "group" {
			var gme messageEvent.GroupMessageEvent
			err = c.BindJSON(&gme)
		}
		if err != nil {
			fmt.Printf("P/G MessageEvent BindJSON error:%s \n", err.Error())
		}
	}
}

func SendTest(c *gin.Context) {
	SendPrivateMessage(utils.CONFIG)
}

func handlePMessage(pme *messageEvent.PrivateMessageEvent) {

}
