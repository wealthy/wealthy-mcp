package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

func placeOrder(ctx context.Context, args falcon.OrderReq) (any, error) {
	return utils.FalconService.PlaceOrder(ctx, []falcon.OrderReq{args})
}

func modifyOrder(ctx context.Context, args falcon.ModifyOrderReq) (any, error) {
	return utils.FalconService.ModifyOrder(ctx, args)
}

func cancelOrder(ctx context.Context, args falcon.CancelOrderReq) (any, error) {
	return utils.FalconService.CancelOrder(ctx, args)
}

func AddOrderTool(mcp *server.MCPServer) {
	PlaceOrderTool.Register(mcp)
	ModifyOrderTool.Register(mcp)
	CancelOrderTool.Register(mcp)
}

var PlaceOrderTool = mcp.MustTool(
	"place_order",
	"Tool for placing buy/sell order",
	placeOrder,
)

var ModifyOrderTool = mcp.MustTool(
	"modify_order",
	"Tool for modifying an order",
	modifyOrder,
)

var CancelOrderTool = mcp.MustTool(
	"cancel_order",
	"Tool for cancelling an order",
	cancelOrder,
)
