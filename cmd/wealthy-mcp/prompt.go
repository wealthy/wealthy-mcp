// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mark3labs/mcp-go/mcp"
)

var (
	placeOrderPromptText = `search for {{company_name}} using "search" tool to get token, exchange_name, trading_symbol if already not present
	get price of {{company_name}} if price type is not market using "price" tool
	call place order with {{token}} {{exchange_name}} {{trading_symbol}} {{quantity}} {{transaction_type}} {{price_type}} {{price}} {{order_source}} {{order_type}}
	`
	getTradeIdeasPromptText = `get trade ideas tool
	call "research" tool
	if "research" tool returns error, search internet for trending stocks to buy
	do analysis of trending stocks to buy before suggesting it to user
	`
	createWatchlistPromptText = `create watchlist
	get watchlist name from user, call get_watchlist tool to get watchlist,
	check if watchlist already exists, if not create watchlist,
	get scrip token, exchage, trading symbol from scrip user search tool,
	call "update_watchlist" tool to add scrip to watchlist
	`
)

func promptHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	args := req.Params.Arguments
	slog.Info("promptrequest", "args", args)
	return &mcp.GetPromptResult{
		Description: "wealthy mcp tool prompt",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("What is the price of %s", args["symbol"]),
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "show me my current holdings",
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "show me my portfolio",
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "show me my positions",
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "I want to buy 100 shares of TATAMOTORS at market price",
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "do SWOT analysis for each stock in my portfolio",
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "show me today's trade ideas",
				},
			},
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: "get price of RELIANCE",
				},
			},
		},
	}, nil
}

func placeOrderPrompt() mcp.Prompt {
	return mcp.NewPrompt("place-order",
		mcp.WithPromptDescription("Place an order for stocks"),
		mcp.WithArgument("trading_symbol", mcp.ArgumentDescription("The trading symbol of the stock (e.g., RELIANCE, TCS)"), mcp.RequiredArgument()),
		mcp.WithArgument("quantity", mcp.ArgumentDescription("Number of shares to buy/sell"), mcp.RequiredArgument()),
		mcp.WithArgument("transaction_type", mcp.ArgumentDescription("1 for Buy, 2 for Sell"), mcp.RequiredArgument()),
		mcp.WithArgument("price_type", mcp.ArgumentDescription("MKT(2) for market price, LMT(1) for limit price"), mcp.RequiredArgument()),
		mcp.WithArgument("price", mcp.ArgumentDescription("Price for limit orders (required if price_type is LMT)")),
		mcp.WithArgument("order_type", mcp.ArgumentDescription("1 for Market, 2 for Limit, 3 for Stop, 4 for Stop Limit"), mcp.RequiredArgument()),
	)
}

func getTradeIdeasPrompt() mcp.Prompt {
	return mcp.NewPrompt("get-trade-ideas",
		mcp.WithPromptDescription("Get trade ideas"),
	)
}

func createWatchlistPrompt() mcp.Prompt {
	return mcp.NewPrompt("create-watchlist",
		mcp.WithPromptDescription("Create a watchlist"),
		mcp.WithArgument("watchlist_name", mcp.ArgumentDescription("Name of the watchlist"), mcp.RequiredArgument()),
		mcp.WithArgument("scrip", mcp.ArgumentDescription("Scrip to add to the watchlist"), mcp.RequiredArgument()),
	)
}

func placeOrderPromptHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	// args := req.Params.Arguments
	// slog.Info("promptrequest for place order", "args", args)
	return &mcp.GetPromptResult{
		Description: "Place an order for stocks",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: placeOrderPromptText,
				},
			},
		},
	}, nil
}

func getTradeIdeasPromptHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Get trade ideas",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: getTradeIdeasPromptText,
				},
			},
		},
	}, nil
}

func createWatchlistPromptHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Create a watchlist",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: createWatchlistPromptText,
				},
			},
		},
	}, nil
}
