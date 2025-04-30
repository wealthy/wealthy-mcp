// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tools

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/server"
	mcp "github.com/wealthy/wealthy-mcp"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
	"go.uber.org/zap"
)

// Default timeout for HTTP client in seconds
const defaultTimeout = 60

var (
	client        = http.Client{Timeout: time.Duration(defaultTimeout) * time.Second}
	falconService = falcon.NewFalconService(&client)
)

// validatePlaceOrderRequest validates the parameters for a place order request
func validatePlaceOrderRequest(args falcon.FalconRequest) error {
	if args.TradingSymbol == "" {
		return fmt.Errorf("trading_symbol is required")
	}
	if args.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	if args.TransactionType != 1 && args.TransactionType != 2 {
		return fmt.Errorf("transaction_type must be 1 (Buy) or 2 (Sell)")
	}
	return nil
}

// validateSecurityInfoRequest validates the parameters for a security info request
func validateSecurityInfoRequest(args falcon.FalconRequest) error {
	if args.TradingSymbol == "" {
		return fmt.Errorf("trading_symbol is required")
	}
	if args.ExchangeName == 0 {
		return fmt.Errorf("exchange_name is required")
	}
	return nil
}

func queryFalcon(ctx context.Context, args falcon.FalconRequest) (any, error) {
	logger, _ := zap.NewProduction(
		zap.WithCaller(true),
		zap.Fields(
			zap.String("query_type", args.QueryType),
			zap.String("trading_symbol", args.TradingSymbol),
		),
	)
	ctx = context.WithValue(ctx, "logger", logger)

	// Add request tracing
	logger.Info("processing falcon request")
	defer logger.Info("completed falcon request")

	switch args.QueryType {
	case "place_order":
		if err := validatePlaceOrderRequest(args); err != nil {
			logger.Error("invalid place order request", zap.Error(err))
			return nil, fmt.Errorf("invalid place order request: %w", err)
		}
		return falconService.PlaceOrder(ctx, args)

	case "get_holdings":
		return falconService.GetHoldings(ctx)

	case "get_positions":
		return falconService.GetPositions(ctx)

	case "get_order_book":
		return falconService.GetOrderBook(ctx)

	case "get_trade_ideas":
		return falconService.GetTradeIdeas(ctx)
	case "get_security_info":
		return falconService.GetSecurityInfo(ctx, &falcon.SecurityInfoReq{
			Name: args.TradingSymbol,
		})
	case "get_price":
		if args.TradingSymbol == "" {
			logger.Error("trading symbol is required for price query")
			return nil, fmt.Errorf("trading_symbol is required for price query")
		}
		return falconService.GetPrice(ctx, &falcon.PriceReq{
			Symbols: []string{"nse:" + args.TradingSymbol},
		})

	default:
		err := fmt.Errorf("unsupported query type: %s", args.QueryType)
		logger.Error("invalid query type", zap.Error(err))
		return nil, err
	}
}

// FalconTool is the MCP tool for interacting with the Falcon trading platform
var FalconTool = mcp.MustTool(
	"falcon_tool",
	"Tool for interacting with the Falcon trading platform",
	queryFalcon,
)

// AddFalconTool registers the Falcon tool with the MCP server
func AddFalconTool(mcp *server.MCPServer) {
	FalconTool.Register(mcp)
}
