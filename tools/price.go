package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

type GetPriceArgs struct {
	Symbols []string `json:"symbols" jsonschema:"description=Symbol of the stock, add -EQ in the end for trading symbol if already not present, correct format: exchange:trading_symbol, nse:RELIANCE-EQ, bse:RELIANCE-EQ, nse:INFY-EQ, bse:INFY, nfo:RELIANCE29MAY25F, 1-nse, 2-nfo, 3-bse, 4-bfo"`
}

func getPrice(ctx context.Context, args GetPriceArgs) (any, error) {
	return utils.FalconService.GetPrice(ctx, &falcon.PriceReq{
		Mode:    3,
		Symbols: args.Symbols,
	})
}

var priceTool = mcp.MustTool(
	"get_price",
	"Get the price of a stock",
	getPrice,
)

func AddPriceTool(mcp *server.MCPServer) {
	priceTool.Register(mcp)
}
