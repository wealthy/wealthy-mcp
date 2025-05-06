package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

type ResearchReq struct {
}

func getResearch(ctx context.Context, args ResearchReq) (any, error) {
	return utils.FalconService.GetTradeIdeas(ctx)
}

var ResearchTool = mcp.MustTool(
	"research",
	"Tool for getting research ideas",
	getResearch,
)

func AddResearchTool(mcp *server.MCPServer) {
	ResearchTool.Register(mcp)
}
