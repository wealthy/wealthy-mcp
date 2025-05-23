package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

type OrderReqArgs struct {
	Order falcon.OrderReq `json:"order" jsonschema:"description=Order to be placed"`
}

func getOrder(ctx context.Context, args falcon.OrderReq) (any, error) {
	return utils.FalconService.PlaceOrder(ctx, []falcon.OrderReq{args})
}

var OrderTool = mcp.MustTool(
	"place_order",
	"Tool for placing buy/sell order",
	getOrder,
)

func AddOrderTool(mcp *server.MCPServer) {
	OrderTool.Register(mcp)
}
