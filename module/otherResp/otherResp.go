package otherResp

type ExchangeRateResp struct {
	Records []ExchangeRate `json:"records"`
}
type ExchangeRate struct {
	Date   string   `json:"date"`
	Values []string `json:"values"`
}
