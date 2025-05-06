package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

func addToWatchlist(ctx context.Context, args falcon.WatchlistReq) (any, error) {
	return utils.FalconService.AddToWatchlist(ctx, &args)
}

var WatchlistTool = mcp.MustTool(
	"create_watchlist",
	"Tool for creating a watchlist",
	addToWatchlist,
)

func AddWatchlistTool(mcp *server.MCPServer) {
	WatchlistTool.Register(mcp)
}

var GetWatchlistTool = mcp.MustTool(
	"get_watchlist",
	"Tool for getting a watchlist",
	getWatchlist,
)

func getWatchlist(ctx context.Context, args falcon.WatchlistReq) (any, error) {
	return utils.FalconService.GetWatchlists(ctx, &args)
}

func AddGetWatchlistTool(mcp *server.MCPServer) {
	GetWatchlistTool.Register(mcp)
}
