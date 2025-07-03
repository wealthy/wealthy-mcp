package websocket

type PriceSubscriptionReq struct {
	Operation int      `json:"operation"`
	Mode      int      `json:"mode"`
	Symbol    []string `json:"symbol"`
}
