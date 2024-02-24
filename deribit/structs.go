package deribit

type requestMessage struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

type tickerMessage struct {
	Method string `json:"method"`
	Params struct {
		Data struct {
			CurrentFunding float64 `json:"current_funding"`
		} `json:"data"`
	} `json:"params"`
}
