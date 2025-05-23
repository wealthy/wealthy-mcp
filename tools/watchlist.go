package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

func addToWatchlist(ctx context.Context, args falcon.WatchlistReq) (any, error) {
	return utils.FalconService.CreateWatchlist(ctx, args.Name)
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

var UpdateWatchlistTool = mcp.MustTool(
	"update_watchlist",
	"Tool for updating a watchlist",
	updateWatchlist,
)

func UpdateWatchlist(mcp *server.MCPServer) {
	UpdateWatchlistTool.Register(mcp)
}

func updateWatchlist(ctx context.Context, args falcon.WatchlistReq) (any, error) {
	return utils.FalconService.AddToWatchlist(ctx, &args)
}

func getWatchlist(ctx context.Context, req getWatchlistReq) (any, error) {
	return utils.FalconService.GetWatchlists(ctx)
}

func AddGetWatchlistTool(mcp *server.MCPServer) {
	GetWatchlistTool.Register(mcp)
}

type getWatchlistReq struct {
}
