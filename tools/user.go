// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

var GetUserMarginTool = mcp.MustTool(
	"get_user_margin",
	"Tool for getting user margin",
	getUserMargin,
)

type GetUserMarginReq struct {
}

func getUserMargin(ctx context.Context, args GetUserMarginReq) (any, error) {
	return utils.FalconService.GetUserMargin(ctx)
}

func AddUserTool(mcp *server.MCPServer) {
	GetUserMarginTool.Register(mcp)
}
