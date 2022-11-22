package service

import (
	"Alpha-S/module/otherResp"
	"Alpha-S/utils"
	"encoding/json"
	"fmt"
)

func GetExchangeRate(msg string) (info string) {
	url := fmt.Sprintf("%s/CcprHisNew?currency=%s&pageNum=1&pageSize=10", utils.CONFIG.OtherApi.ExchangeRate, msg)
	bytes, err := utils.SendHttpRequest("GET", url, nil, nil)
	if err != nil {
		fmt.Printf("GetExchangeRate error:%s \n", err.Error())
		return
	}
	var resp otherResp.ExchangeRateResp
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		fmt.Printf("GetExchangeRate Unmarshal error:%s \n", err.Error())
		return
	}

	info = fmt.Sprintf("更新汇率成功!\n今日%s汇率：%s", msg, resp.Records[0])
	fmt.Printf("实时汇率:%s", info)
	return

	return
}
