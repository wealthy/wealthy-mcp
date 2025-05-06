package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

type SearchReq struct {
	Query string `json:"query" jsonschema:"description=Search query for finding a security symbol"`
}

func getSearch(ctx context.Context, args SearchReq) (any, error) {
	return utils.FalconService.GetSecurityInfo(ctx, &falcon.SecurityInfoReq{
		Name: args.Query,
	})
}

var SearchTool = mcp.MustTool(
	"search",
	"Tool for searching for a symbol",
	getSearch,
)

func AddSearchTool(mcp *server.MCPServer) {
	SearchTool.Register(mcp)
}
