// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

var (
	placeOrderPromptText = ` Place an order for stocks
		1. If the token, exchange name, or trading symbol for {{trading_symbol}} is not already available:
   			- Use the "search" tool to retrieve the token, exchange name, and trading symbol.
		2. If the {{price_type}} is not "market":
   			- Use the "price" tool to get the latest price for {{trading_symbol}}.
		3. calculate required margin for the order using "get_user_margin" tool
		4. if required margin is greater than user margin, then ask user to add more funds using wealthy app
		5. Call the "place_order" tool with the following parameters:
			- token: {{token}}
			- exchange_name: {{exchange_name}}
			- trading_symbol: {{trading_symbol}}
			- quantity: {{quantity}}
			- transaction_type: {{transaction_type}} (e.g., Buy or Sell)
			- price_type: {{price_type}} (e.g., Market or Limit)
			- price: {{price}} (required only for limit orders)
			- order_source: {{order_source}}
			- order_type: {{order_type}} (e.g., CNC, MIS, etc.)`

	getTradeIdeasPromptText = ` Get trade ideas
		1. Call the "research" tool to get trade ideas.
	2. If the "research" tool returns an error, search the internet for trending stocks to buy.
	3. Perform an analysis of the trending stocks to buy before suggesting them to the user.
	`
	createWatchlistPromptText = ` Create a watchlist
		1. Ask the user for a name for the new watchlist.
		2. Use the "get_watchlist" tool to check if a watchlist with that name already exists.
		3. If the watchlist does not exist, create a new watchlist with the provided name.
		4. Ask the user to search for a stock (scrip) to add.
		5. Use the "scrip_user_search" tool to get the scrip token, exchange, and trading symbol based on the user's input.
		6. Use the "update_watchlist" tool to add the selected scrip to the specified watchlist.
	`
	portfolioAnalysisPromptText = `
		Perform portfolio analysis of user holdings
		1. Use the "report" tool to retrieve the user's current portfolio holdings.
		2. For each holding:
   			- Perform a SWOT (Strengths, Weaknesses, Opportunities, Threats) analysis using up-to-date internet search.
   			- Assign a rating to each stock based on the SWOT analysis (e.g., Strong Buy, Buy, Hold, Sell, Strong Sell).
		3. Use the "get_price" tool to retrieve the latest prices for all stocks in the portfolio to calculate the total value of the portfolio. Multiple stock symbols may be passed at once. if ltp is zero then check internet for latest price of the stock in NSE exchange
		4. Summarize the portfolio analysis with a brief overview highlighting strengths, weaknesses, and key insights.
`
)

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

func portfolioAnalysisPrompt() mcp.Prompt {
	return mcp.NewPrompt("portfolio-analysis",
		mcp.WithPromptDescription("Performs portfolio analysis of user holdings"),
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

func portfolioAnalysisPromptHandler(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Portfolio analysis",
		Messages: []mcp.PromptMessage{
			{
				Role: mcp.RoleUser,
				Content: mcp.TextContent{
					Type: "text",
					Text: portfolioAnalysisPromptText,
				},
			},
		},
	}, nil
}
