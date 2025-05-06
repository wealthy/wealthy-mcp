package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/utils"
)

const (
	ReportTypeHoldings  = "holdings"
	ReportTypePositions = "positions"
	ReportTypeOrders    = "order_book"
)

type ReportRequest struct {
	Report string `json:"report" jsonschema:"description=Report type, holdings=holdings, positions=positions, order_book=order_book"`
}

func getReports(ctx context.Context, args ReportRequest) (any, error) {
	switch args.Report {
	case ReportTypeHoldings:
		return utils.FalconService.GetHoldings(ctx)
	case ReportTypePositions:
		return utils.FalconService.GetPositions(ctx)
	case ReportTypeOrders:
		return utils.FalconService.GetOrderBook(ctx)
	default:
		return nil, fmt.Errorf("unsupported report type: %s", args.Report)
	}
}

var ReportsTool = mcp.MustTool(
	"reports_tool",
	"Tool for generating reports",
	getReports,
)

func AddReportsTool(mcp *server.MCPServer) {
	ReportsTool.Register(mcp)
}
