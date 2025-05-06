package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

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

func getOrder(ctx context.Context, args OrderReq) (any, error) {
	return utils.FalconService.PlaceOrder(ctx, falcon.FalconRequest{
		QueryType: "place_order",
		OrderReq:  falcon.OrderReq(args),
	})
}

var OrderTool = mcp.MustTool(
	"place_order",
	"Tool for placing an order",
	getOrder,
)

func AddOrderTool(mcp *server.MCPServer) {
	OrderTool.Register(mcp)
}
