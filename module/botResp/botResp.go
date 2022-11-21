package botResp

type BotResp struct {
	Status  string         `json:"status"`
	RetCode int32          `json:"retcode"`
	Msg     string         `json:"msg"`
	Wording string         `json:"wording"`
	Data    any            `json:"data"`
	Echo    map[string]any `json:"echo"`
}
