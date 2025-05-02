// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package falcon

import "time"

type FalconRequest struct {
	QueryType string `json:"query_type" jsonschema:"description=Type of query (place_order/get_holdings/get_positions/get_security_info/get_order_book/get_price(prices are in paisa)/get_trade_ideas)"`
	OrderReq
	SecurityInfoReq
}

type OrderReq struct {
	ExchangeName    int    `json:"exchange_name" jsonschema:"description=Exchange name identifier, NSE=1, NFO=2, BSE=3, BFO=4"`
	Token           string `json:"token" jsonschema:"description=Trading token"`
	TradingSymbol   string `json:"trading_symbol" jsonschema:"description=Symbol to trade"`
	Quantity        int    `json:"quantity" jsonschema:"description=Quantity to trade"`
	Price           string `json:"price" jsonschema:"description=Price for the order"`
	TriggerPrice    string `json:"trigger_price,omitempty" jsonschema:"description=Trigger price for stop orders"`
	OrderType       int    `json:"order_type" jsonschema:"description=Type of order, 1=Market, 2=Limit, 3=Stop, 4=Stop Limit"`
	TransactionType int    `json:"transaction_type" jsonschema:"description=Buy (1) or Sell (2)"`
	PriceType       int    `json:"price_type" jsonschema:"description=Price type for the order, 1=LMT (Limit), 2=MKT (Market), 3=SLLMT (Stop Loss Limit), 4=SLMKT (Stop Loss Market), 5=DS (Disclosed), 6=TWOLEG (Two Leg), 7=THREEELEG (Three Leg)"`
	Validity        int    `json:"validity" jsonschema:"description=Validity of the order"`
	DiscQuantity    int    `json:"disclosed_quantity" jsonschema:"description=Disclosed quantity for the order"`
	IsAMO           bool   `json:"is_amo" jsonschema:"description=Whether this is an After Market Order"`

	// Protection parameters
	TargetPrice   string `json:"target_price,omitempty" jsonschema:"description=Target price for the order"`
	StopLossPrice string `json:"stop_loss_price,omitempty" jsonschema:"description=Stop loss price for the order"`
	TrailPrice    string `json:"trailing_price,omitempty" jsonschema:"description=Trailing price for the order"`
}

type Order struct {
	UserID          string `json:"-"`
	ExchangeOrderID string `json:"exchange_order_id,omitempty"`
	ParentOrderID   string `json:"parent_order_id,omitempty"`
	// IsAMO                 bool   `json:"-"`
	RejectReason          string `json:"reject_reason,omitempty"`
	MarketProdPercentange string `json:"-"`
	Status                int    `json:"status"`
	OriginalQty           int    `json:"original_quantity,omitempty"`
	OriginalPrice         string `json:"original_price,omitempty"`
	FilledShares          int    `json:"filled_shares,omitempty"`
	CancelledOrderQty     int    `json:"cancelled_quantity,omitempty"`
	AvgPrice              string `json:"average_price,omitempty"`
	PricePrecision        string `json:"price_precision,omitempty"`
	LotSize               int    `json:"lot_size,omitempty"`
	TickSize              string `json:"tick_size,omitempty"`
	PriceFactor           string `json:"price_factor,omitempty"`
	OrderEntryTime        string `json:"entry_time"`
	OmsTime               string `json:"oms_time"`
	ExchangeTime          string `json:"exchange_time,omitempty"`
	TrgtStoplossHit       string `json:"target_stoploss_hit,omitempty"`
	RTargetPrice          string `json:"req_target_price,omitempty"`
	OrgTargetPrice        string `json:"original_target_price,omitempty"`
	RStoplossPrice        string `json:"req_stop_loss_price,omitempty"`
	OrgStopLossPrice      string `json:"org_stop_loss_price,omitempty"`
	Tags                  string `json:"tags,omitempty"` //comma separated values
	FillTime              string `json:"fill_time,omitempty"`
	FillID                string `json:"fill_id,omitempty"`
	FillQty               string `json:"fill_quantity,omitempty"`
	FillPrice             string `json:"fill_price,omitempty"`
	ReportType            int    `json:"report_type,omitempty"`
	BasketOrderID         string `json:"basket_order_id,omitempty"`
	ClientID              string `json:"-"`
	// RejectedBy            string `json:"rejected_by,omitempty"`
	UserOrder

	//	Legs                  []Leg       `json:"legs"` //not required for current release
	//Tradebook specific fields
	// CstFirm   string `json:"customer_firm"`
}

type ProtectioParam struct {
	TargetPrice   string `json:"target_price,omitempty"`
	StopLossPrice string `json:"stop_loss_price,omitempty"`
	TrailPrice    string `json:"trailing_price,omitempty"`
}

type Leg struct {
	TradingSymbol   string `json:"trading_symbol"`
	Quantity        int    `json:"quantity"`
	Price           string `json:"price"`
	TransactionType int    `json:"transaction_type"`
}

type UserOrder struct {
	OrderSource int `json:"order_source"`
	ModifyOrder
}

type ModifyOrder struct {
	OMSID   string `json:"oms_id"`
	OrderID string `json:"order_id"`
	Remarks string `json:"-"`
	ProtectioParam
	OrderItem
}

type Margin struct {
	Cash                 string `json:"cash,omitempty"`
	MarginUsed           string `json:"margin_used"`
	MarginUsedAfterTrade string `json:"margin_used_after_trade,omitempty"`
	OrderMargin          string `json:"order_margin,omitempty"`
	PrevMarginUsed       string `json:"margin_used_previous,omitempty"`
	Remarks              string `json:"remarks,omitempty"`
}

type OrderItem struct {
	ExchangeName    int       `json:"exchange_name"`
	Token           string    `json:"token"`
	TradingSymbol   string    `json:"trading_symbol"`
	Quantity        int       `json:"quantity"`
	Price           string    `json:"price"`
	TriggerPrice    string    `json:"trigger_price,omitempty"`
	OrderType       int       `json:"order_type"`
	TransactionType int       `json:"transaction_type"`
	PriceType       int       `json:"price_type"`
	Validity        int       `json:"validity"`
	DiscQuantity    int       `json:"disclosed_quantity"`
	IsAMO           bool      `json:"is_amo"`
	CreatedAt       time.Time `json:"-" `
	UpdatedAt       time.Time `json:"-"`
}

type SecurityInfoReq struct {
	Name string `json:"name" jsonschema:"required,description=Symbol or name for a security"`
}

type PriceReq struct {
	Mode    int      `json:"mode"`
	Symbols []string `json:"symbols"`
}
