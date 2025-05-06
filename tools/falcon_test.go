// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tools

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wealthy/wealthy-mcp/internal/falcon"
)

func TestValidatePlaceOrderRequest(t *testing.T) {
	tests := []struct {
		name    string
		args    falcon.FalconRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        100,
					TransactionType: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "missing account ID",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        100,
					TransactionType: 1,
				},
			},
			wantErr: true,
			errMsg:  "account_id is required",
		},
		{
			name: "missing trading symbol",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        100,
					TransactionType: 1,
				},
			},
			wantErr: true,
			errMsg:  "trading_symbol is required",
		},
		{
			name: "invalid quantity",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        0,
					TransactionType: 1,
				},
			},
			wantErr: true,
			errMsg:  "quantity must be greater than 0",
		},
		{
			name: "invalid transaction type",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        100,
					TransactionType: 3,
				},
			},
			wantErr: true,
			errMsg:  "transaction_type must be 1 (Buy) or 2 (Sell)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePlaceOrderRequest(tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateSecurityInfoRequest(t *testing.T) {
	tests := []struct {
		name    string
		args    falcon.FalconRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			args: falcon.FalconRequest{
				QueryType: "get_security_info",
				SecurityInfoReq: falcon.SecurityInfoReq{
					Name: "AAPL",
				},
			},
			wantErr: false,
		},
		{
			name: "missing trading symbol",
			args: falcon.FalconRequest{
				QueryType: "get_security_info",
				OrderReq: falcon.OrderReq{
					TradingSymbol: "AAPL",
					ExchangeName:  1,
				},
			},
			wantErr: true,
			errMsg:  "trading_symbol is required",
		},
		{
			name: "missing exchange name",
			args: falcon.FalconRequest{
				QueryType: "get_security_info",
				OrderReq: falcon.OrderReq{
					TradingSymbol: "AAPL",
				},
			},
			wantErr: true,
			errMsg:  "exchange_name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateSecurityInfoRequest(tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// MockFalconService is a mock implementation of the Falcon service for testing
type MockFalconService struct {
	PlaceOrderFunc      func(ctx context.Context, req falcon.FalconRequest) (*falcon.Order, error)
	GetHoldingsFunc     func(ctx context.Context) (any, error)
	GetPositionsFunc    func(ctx context.Context) (any, error)
	GetOrderBookFunc    func(ctx context.Context) (any, error)
	GetTradeIdeasFunc   func(ctx context.Context) (any, error)
	GetSecurityInfoFunc func(ctx context.Context, req *falcon.SecurityInfoReq) (any, error)
	GetPriceFunc        func(ctx context.Context, req *falcon.PriceReq) (any, error)
	AddToWatchlistFunc  func(ctx context.Context, req *falcon.WatchlistReq) (any, error)
	GetWatchlistsFunc   func(ctx context.Context, req *falcon.WatchlistReq) (any, error)
}

func (m *MockFalconService) PlaceOrder(ctx context.Context, req falcon.FalconRequest) (*falcon.Order, error) {
	return m.PlaceOrderFunc(ctx, req)
}

func (m *MockFalconService) GetHoldings(ctx context.Context) (any, error) {
	return m.GetHoldingsFunc(ctx)
}

func (m *MockFalconService) GetPositions(ctx context.Context) (any, error) {
	return m.GetPositionsFunc(ctx)
}

func (m *MockFalconService) GetOrderBook(ctx context.Context) (any, error) {
	return m.GetOrderBookFunc(ctx)
}

func (m *MockFalconService) GetTradeIdeas(ctx context.Context) (any, error) {
	return m.GetTradeIdeasFunc(ctx)
}

func (m *MockFalconService) GetSecurityInfo(ctx context.Context, req *falcon.SecurityInfoReq) (any, error) {
	return m.GetSecurityInfoFunc(ctx, req)
}

func (m *MockFalconService) GetPrice(ctx context.Context, req *falcon.PriceReq) (any, error) {
	return m.GetPriceFunc(ctx, req)
}

func (m *MockFalconService) AddToWatchlist(ctx context.Context, req *falcon.WatchlistReq) (any, error) {
	return m.AddToWatchlistFunc(ctx, req)
}

func (m *MockFalconService) GetWatchlists(ctx context.Context, req *falcon.WatchlistReq) (any, error) {
	return m.GetWatchlistsFunc(ctx, req)
}

func TestQueryFalconWithMocks(t *testing.T) {
	// Save original service and restore after test
	originalService := falconService
	defer func() { falconService = originalService }()

	mockService := &MockFalconService{
		PlaceOrderFunc: func(ctx context.Context, req falcon.FalconRequest) (*falcon.Order, error) {
			return &falcon.Order{
				UserOrder: falcon.UserOrder{
					ModifyOrder: falcon.ModifyOrder{
						OMSID: "123",
					},
				},
			}, nil
		},
		GetHoldingsFunc: func(ctx context.Context) (any, error) {
			return []map[string]interface{}{{"symbol": "AAPL", "quantity": 100}}, nil
		},
		GetPositionsFunc: func(ctx context.Context) (any, error) {
			return []map[string]interface{}{{"symbol": "GOOGL", "quantity": 50}}, nil
		},
		GetOrderBookFunc: func(ctx context.Context) (any, error) {
			return []map[string]interface{}{{"order_id": "123", "status": "COMPLETE"}}, nil
		},
		GetTradeIdeasFunc: func(ctx context.Context) (any, error) {
			return []map[string]interface{}{{"symbol": "MSFT", "signal": "BUY"}}, nil
		},
		GetSecurityInfoFunc: func(ctx context.Context, req *falcon.SecurityInfoReq) (any, error) {
			return map[string]interface{}{"name": req.Name, "exchange": "NSE"}, nil
		},
		GetPriceFunc: func(ctx context.Context, req *falcon.PriceReq) (any, error) {
			return map[string]interface{}{"price": 150.50}, nil
		},
		AddToWatchlistFunc: func(ctx context.Context, req *falcon.WatchlistReq) (any, error) {
			return map[string]interface{}{"status": "success"}, nil
		},
		GetWatchlistsFunc: func(ctx context.Context, req *falcon.WatchlistReq) (any, error) {
			return []map[string]interface{}{{"name": "Default", "symbols": []string{"AAPL", "GOOGL"}}}, nil
		},
	}

	falconService = mockService

	tests := []struct {
		name    string
		args    falcon.FalconRequest
		want    any
		wantErr bool
	}{
		{
			name: "successful place order",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        100,
					TransactionType: 1,
				},
			},
			want: &falcon.Order{
				UserOrder: falcon.UserOrder{
					ModifyOrder: falcon.ModifyOrder{
						OMSID: "123",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "successful get holdings",
			args: falcon.FalconRequest{
				QueryType: "get_holdings",
			},
			want:    []map[string]interface{}{{"symbol": "AAPL", "quantity": 100}},
			wantErr: false,
		},
		{
			name: "successful get positions",
			args: falcon.FalconRequest{
				QueryType: "get_positions",
			},
			want:    []map[string]interface{}{{"symbol": "GOOGL", "quantity": 50}},
			wantErr: false,
		},
		{
			name: "successful get order book",
			args: falcon.FalconRequest{
				QueryType: "get_order_book",
			},
			want:    []map[string]interface{}{{"order_id": "123", "status": "COMPLETE"}},
			wantErr: false,
		},
		{
			name: "successful get trade ideas",
			args: falcon.FalconRequest{
				QueryType: "get_trade_ideas",
			},
			want:    []map[string]interface{}{{"symbol": "MSFT", "signal": "BUY"}},
			wantErr: false,
		},
		{
			name: "successful get security info",
			args: falcon.FalconRequest{
				QueryType: "get_security_info",
				SecurityInfoReq: falcon.SecurityInfoReq{
					Name: "AAPL",
				},
			},
			want:    map[string]interface{}{"name": "AAPL", "exchange": "NSE"},
			wantErr: false,
		},
		{
			name: "successful get price",
			args: falcon.FalconRequest{
				QueryType: "get_price",
				OrderReq: falcon.OrderReq{
					TradingSymbol: "AAPL",
				},
			},
			want:    map[string]interface{}{"price": 150.50},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := queryFalcon(ctx, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestQueryFalconErrors(t *testing.T) {
	tests := []struct {
		name    string
		args    falcon.FalconRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "invalid query type",
			args: falcon.FalconRequest{
				QueryType: "invalid_type",
			},
			wantErr: true,
			errMsg:  "unsupported query type: invalid_type",
		},
		{
			name: "invalid place order request - missing trading symbol",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					Quantity:        100,
					TransactionType: 1,
				},
			},
			wantErr: true,
			errMsg:  "invalid place order request: trading_symbol is required",
		},
		{
			name: "invalid place order request - invalid quantity",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        0,
					TransactionType: 1,
				},
			},
			wantErr: true,
			errMsg:  "invalid place order request: quantity must be greater than 0",
		},
		{
			name: "invalid place order request - invalid transaction type",
			args: falcon.FalconRequest{
				QueryType: "place_order",
				OrderReq: falcon.OrderReq{
					TradingSymbol:   "AAPL",
					Quantity:        100,
					TransactionType: 3,
				},
			},
			wantErr: true,
			errMsg:  "invalid place order request: transaction_type must be 1 (Buy) or 2 (Sell)",
		},
		{
			name: "invalid get price request - missing trading symbol",
			args: falcon.FalconRequest{
				QueryType: "get_price",
			},
			wantErr: true,
			errMsg:  "trading_symbol is required for price query",
		},
		{
			name: "invalid security info request - missing name",
			args: falcon.FalconRequest{
				QueryType: "get_security_info",
			},
			wantErr: true,
			errMsg:  "invalid security info request: search query is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := queryFalcon(ctx, tt.args)
			assert.Error(t, err)
			assert.Equal(t, tt.errMsg, err.Error())
		})
	}
}
