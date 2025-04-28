// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package falcon

import "time"

type FalconRequest struct {
	// Order placement parameters
	// AccessToken     string `json:"access_token" jsonschema:"description=Access token for the user"`
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

	// Query parameters
	QueryType string `json:"query_type" jsonschema:"description=Type of query (place_order/get_holdings/get_positions/get_security_info/get_order_book)"`
}

type Order struct {
	UserID          string `json:"-" db:"user_id"`
	ExchangeOrderID string `json:"exchange_order_id,omitempty" db:"exchange_order_id"`
	ParentOrderID   string `json:"parent_order_id,omitempty" db:"parent_order_id"`
	// IsAMO                 bool   `json:"-" db:"is_amo"`
	RejectReason          string `json:"reject_reason,omitempty" db:"reject_reason"`
	MarketProdPercentange string `json:"-" db:"-"`
	Status                int    `json:"status" db:"status"`
	OriginalQty           int    `json:"original_quantity,omitempty" db:"original_quantity"`
	OriginalPrice         string `json:"original_price,omitempty" db:"original_price"`
	FilledShares          int    `json:"filled_shares,omitempty" db:"filled_shares"`
	CancelledOrderQty     int    `json:"cancelled_quantity,omitempty" db:"cancelled_quantity"`
	AvgPrice              string `json:"average_price,omitempty" db:"average_price"`
	PricePrecision        string `json:"price_precision,omitempty" db:"-"`
	LotSize               int    `json:"lot_size,omitempty" db:"lot_size"`
	TickSize              string `json:"tick_size,omitempty" db:"-"`
	PriceFactor           string `json:"price_factor,omitempty" db:"-"`
	OrderEntryTime        string `json:"entry_time" db:"entry_time"`
	OmsTime               string `json:"oms_time" db:"oms_time"`
	ExchangeTime          string `json:"exchange_time,omitempty" db:"exchange_time"`
	TrgtStoplossHit       string `json:"target_stoploss_hit,omitempty" db:"target_stoploss_hit"`
	RTargetPrice          string `json:"req_target_price,omitempty" db:"req_target_price"`
	OrgTargetPrice        string `json:"original_target_price,omitempty" db:"original_target_price"`
	RStoplossPrice        string `json:"req_stop_loss_price,omitempty" db:"req_stop_loss_price"`
	OrgStopLossPrice      string `json:"org_stop_loss_price,omitempty" db:"org_stop_loss_price"`
	Tags                  string `json:"tags,omitempty" db:"tags"` //comma separated values
	FillTime              string `json:"fill_time,omitempty" db:"fill_time"`
	FillID                string `json:"fill_id,omitempty" db:"fill_id"`
	FillQty               string `json:"fill_quantity,omitempty" db:"fill_quantity"`
	FillPrice             string `json:"fill_price,omitempty" db:"fill_price"`
	ReportType            int    `json:"report_type,omitempty" db:"report_type"`
	BasketOrderID         string `json:"basket_order_id,omitempty" db:"basket_order_id"`
	ClientID              string `json:"-" db:"client_id"`
	// RejectedBy            string `json:"rejected_by,omitempty" db:"rejected_by"`
	UserOrder

	//	Legs                  []Leg       `json:"legs" db:""` //not required for current release
	//Tradebook specific fields
	// CstFirm   string `json:"customer_firm" db:""`
}

type ProtectioParam struct {
	TargetPrice   string `json:"target_price,omitempty" db:"target_price"`
	StopLossPrice string `json:"stop_loss_price,omitempty" db:"stop_loss_price"`
	TrailPrice    string `json:"trailing_price,omitempty" db:"trailing_price"`
}

type Leg struct {
	TradingSymbol   string `json:"trading_symbol" db:"trading_symbol"`
	Quantity        int    `json:"quantity" db:"quantity"`
	Price           string `json:"price" db:"price"`
	TransactionType int    `json:"transaction_type" db:"transaction_type"`
}

type UserOrder struct {
	OrderSource int `json:"order_source" db:"order_source"`
	ModifyOrder
}

type ModifyOrder struct {
	OMSID   string `json:"oms_id" db:"oms_id"`
	OrderID string `json:"order_id" db:"order_id"`
	Remarks string `json:"-" db:"remarks"`
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
	ExchangeName    int       `json:"exchange_name" db:"exchange_name"`
	Token           string    `json:"token" db:"token"`
	TradingSymbol   string    `json:"trading_symbol" db:"trading_symbol"`
	Quantity        int       `json:"quantity" db:"quantity"`
	Price           string    `json:"price" db:"price"`
	TriggerPrice    string    `json:"trigger_price,omitempty" db:"trigger_price"`
	OrderType       int       `json:"order_type" db:"order_type"`
	TransactionType int       `json:"transaction_type" db:"transaction_type"`
	PriceType       int       `json:"price_type" db:"price_type"`
	Validity        int       `json:"validity" db:"validity"`
	DiscQuantity    int       `json:"disclosed_quantity" db:"disclosed_quantity"`
	IsAMO           bool      `json:"is_amo" db:"is_amo"`
	CreatedAt       time.Time `json:"-"  db:"created_at"`
	UpdatedAt       time.Time `json:"-" db:"updated_at"`
}

type SecurityInfoReq struct {
	Token        string `json:"symbol"`
	ExchangeName int    `json:"exchange_name"`
}

type PriceReq struct {
	Mode    int      `json:"mode"`
	Symbols []string `json:"symbols"`
}
