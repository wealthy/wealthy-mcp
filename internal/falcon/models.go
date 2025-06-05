// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package falcon

import (
	"strings"
	"time"
)

type FalconRequest struct {
	QueryType string `json:"query_type" jsonschema:"description=Type of query (place_order/get_holdings/get_positions/get_security_info/get_order_book/get_price(prices are in paisa)/get_trade_ideas)"`
	OrderReq
	SecurityInfoReq
}

type BasketOrderReq struct {
	Orders []OrderReq `json:"orders" jsonschema:"description=List of orders to be placed"`
}

// OrderReq defines the schema for placing an order via MCP Falcon.
//
// Field rules:
// - order_type: 1=Market, 2=Limit, 3=Stop, 4=Stop Limit
// - price_type: 1=LMT (Limit), 2=MKT (Market), 3=SLLMT (Stop Loss Limit), 4=SLMKT (Stop Loss Market), 5=DS (Disclosed), 6=TWOLEG (Two Leg), 7=THREEELEG (Three Leg)
// - price: REQUIRED for Limit (order_type=2) and Stop Limit (order_type=4) orders; OMIT or set to empty for Market orders
// - trigger_price: REQUIRED for Stop and Stop Limit orders
// - transaction_type: 1=Buy, 2=Sell
//
// See field-level comments for more details.
type OrderReq struct {
	// Exchange name identifier. NSE=1, NFO=2, BSE=3, BFO=4
	ExchangeName int `json:"exchange_name" jsonschema:"description=Exchange name identifier, NSE=1, NFO=2, BSE=3, BFO=4"`
	// Trading token for the security
	Token string `json:"token" jsonschema:"description=Trading token"`
	// Symbol to trade (e.g., IOC-EQ)
	TradingSymbol string `json:"trading_symbol" jsonschema:"description=Symbol to trade"`
	// Quantity to trade
	Quantity int `json:"quantity" jsonschema:"description=Quantity to trade"`
	// Price for the order. REQUIRED for Limit (order_type=2) and Stop Limit (order_type=4) orders. OMIT or set to empty for Market orders.
	Price string `json:"price" jsonschema:"description=Price for the order. Required for LMT/SL LMT orders."`
	// Trigger price for stop orders. REQUIRED for Stop (order_type=3) and Stop Limit (order_type=4) orders.
	TriggerPrice string `json:"trigger_price,omitempty" jsonschema:"description=Trigger price for stop orders"`
	// Type of order. 1=Market, 2=Limit, 3=Stop, 4=Stop Limit
	OrderType int `json:"order_type" jsonschema:"description=Type of order, 1=Market, 2=Limit, 3=Stop, 4=Stop Limit"`
	// Transaction type. 1=Buy, 2=Sell
	TransactionType int `json:"transaction_type" jsonschema:"description=Buy (1) or Sell (2)"`
	// Price type for the order. 1=LMT (Limit), 2=MKT (Market), 3=SLLMT (Stop Loss Limit), 4=SLMKT (Stop Loss Market), 5=DS (Disclosed), 6=TWOLEG (Two Leg), 7=THREEELEG (Three Leg)
	PriceType int `json:"price_type" jsonschema:"description=Price type for the order, 1=LMT (Limit), 2=MKT (Market), 3=SLLMT (Stop Loss Limit), 4=SLMKT (Stop Loss Market), 5=DS (Disclosed), 6=TWOLEG (Two Leg), 7=THREEELEG (Three Leg)"`
	// Validity of the order (e.g., 1=DAY, 2=IOC, 3=EOS, 4=GTT)
	Validity int `json:"validity" jsonschema:"description=Validity of the order, 1=DAY, 2=IOC, 3=EOS, 4=GTT"`
	// Disclosed quantity for the order
	DiscQuantity int `json:"disclosed_quantity" jsonschema:"description=Disclosed quantity for the order"`
	// Whether this is an After Market Order
	IsAMO bool `json:"is_amo" jsonschema:"description=Whether this is an After Market Order"`

	// Protection parameters
	// Target price for the order (optional)
	TargetPrice string `json:"target_price,omitempty" jsonschema:"description=Target price for the order"`
	// Stop loss price for the order (optional)
	StopLossPrice string `json:"stop_loss_price,omitempty" jsonschema:"description=Stop loss price for the order"`
	// Trailing price for the order (optional)
	TrailPrice string `json:"trailing_price,omitempty" jsonschema:"description=Trailing price for the order"`
	// Order source identifier, always 5
	OrderSource int `json:"order_source" jsonschema:"description=Order source identifier, always 5"`
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
	OrderSource int `json:"order_source" jsonschema:"enum=5,description=Order source identifier, always 5"`
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

type WebsocketURLResponse struct {
	BaseURL string `json:"base_url"`
	Body    any    `json:"body"`
}

type WatchlistReq struct {
	Name   string  `json:"name" jsonschema:"required,description=Name of the watchlist, pass 'All' to get all watchlists"`
	Scrips []scrip `json:"instrument,omitempty" jsonschema:"description=Scrips to add to the watchlist"`
}

type scrip struct {
	Exchange int    `json:"exchange" jsonschema:"description=Exchange name identifier, NSE=1, NFO=2, BSE=3, BFO=4, check security info result"`
	Token    string `json:"token" jsonschema:"description=token of the scrip, check security info result"`
}

func MakePriceReq(symbols []string) *PriceReq {
	priceReq := &PriceReq{
		Mode: 3,
	}
	for _, symbol := range symbols {
		if !strings.HasSuffix(symbol, "-EQ") {
			symbol = symbol + "-EQ"
		}
		priceReq.Symbols = append(priceReq.Symbols, "nse:"+symbol)
	}
	return priceReq
}

type placeOrderResponse struct {
	OrderID       string `json:"order_id"`
	TradingSymbol string `json:"trading_symbol"`
	Quantity      int    `json:"quantity"`
	IsAMO         bool   `json:"is_amo"`
	Status        string `json:"status"`
}

type ModifyOrderReq struct {
	OrderID string `json:"order_id"`
	OrderReq
}

type CancelOrderReq struct {
	OrderType int    `json:"order_type"`
	OrderID   string `json:"order_id"`
}
