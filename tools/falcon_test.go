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
				AccountID:       "ACC123",
				TradingSymbol:   "AAPL",
				Quantity:        100,
				TransactionType: 1,
			},
			wantErr: false,
		},
		{
			name: "missing account ID",
			args: falcon.FalconRequest{
				TradingSymbol:   "AAPL",
				Quantity:        100,
				TransactionType: 1,
			},
			wantErr: true,
			errMsg:  "account_id is required",
		},
		{
			name: "missing trading symbol",
			args: falcon.FalconRequest{
				AccountID:       "ACC123",
				Quantity:        100,
				TransactionType: 1,
			},
			wantErr: true,
			errMsg:  "trading_symbol is required",
		},
		{
			name: "invalid quantity",
			args: falcon.FalconRequest{
				AccountID:       "ACC123",
				TradingSymbol:   "AAPL",
				Quantity:        0,
				TransactionType: 1,
			},
			wantErr: true,
			errMsg:  "quantity must be greater than 0",
		},
		{
			name: "invalid transaction type",
			args: falcon.FalconRequest{
				AccountID:       "ACC123",
				TradingSymbol:   "AAPL",
				Quantity:        100,
				TransactionType: 3,
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
				TradingSymbol: "AAPL",
				ExchangeName:  1,
			},
			wantErr: false,
		},
		{
			name: "missing trading symbol",
			args: falcon.FalconRequest{
				ExchangeName: 1,
			},
			wantErr: true,
			errMsg:  "trading_symbol is required",
		},
		{
			name: "missing exchange name",
			args: falcon.FalconRequest{
				TradingSymbol: "AAPL",
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

func TestValidateAccountRequest(t *testing.T) {
	tests := []struct {
		name      string
		accountID string
		wantErr   bool
		errMsg    string
	}{
		{
			name:      "valid account ID",
			accountID: "ACC123",
			wantErr:   false,
		},
		{
			name:      "empty account ID",
			accountID: "",
			wantErr:   true,
			errMsg:    "account_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAccountRequest(tt.accountID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestQueryFalcon(t *testing.T) {
	tests := []struct {
		name     string
		args     falcon.FalconRequest
		wantErr  bool
		errCheck func(t *testing.T, err error)
	}{
		{
			name: "invalid query type",
			args: falcon.FalconRequest{
				QueryType: "invalid_type",
			},
			wantErr: true,
			errCheck: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "unsupported query type")
			},
		},
		{
			name: "invalid place order request",
			args: falcon.FalconRequest{
				QueryType: "place_order",
			},
			wantErr: true,
			errCheck: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "invalid place order request")
			},
		},
		{
			name: "invalid get holdings request",
			args: falcon.FalconRequest{
				QueryType: "get_holdings",
			},
			wantErr: true,
			errCheck: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "invalid get holdings request")
			},
		},
		{
			name: "invalid get price request",
			args: falcon.FalconRequest{
				QueryType: "get_price",
			},
			wantErr: true,
			errCheck: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "trading_symbol is required")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := queryFalcon(ctx, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errCheck != nil {
					tt.errCheck(t, err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
