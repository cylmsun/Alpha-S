package otherResp

type ExchangeRateResp struct {
	Records []ExchangeRate `json:"records"`
	Data    ExchangeData   `json:"data"`
}
type ExchangeRate struct {
	Date   string   `json:"date"`
	Values []string `json:"values"`
}
type ExchangeData struct {
	Currency string `json:"currency"`
}
