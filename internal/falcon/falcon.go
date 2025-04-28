// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package falcon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wealthy/go-kit/web"
	"github.com/wealthy/wealthy-mcp/internal"
)

var (
	falconBaseURL = "https://api.wealthy.in/broking/api"
	midasBaseURL  = "https://api.wealthy.in/midas/api"
)

// FalconRequest represents the common parameters for Falcon API requests

// FalconService defines the interface for Falcon API operations
type FalconService interface {
	PlaceOrder(ctx context.Context, req FalconRequest) (*Order, error)
	GetHoldings(ctx context.Context, accountID string) (any, error)
	GetPositions(ctx context.Context, accountID string) (any, error)
	GetSecurityInfo(ctx context.Context, req *SecurityInfoReq) (any, error)
	GetOrderBook(ctx context.Context) (any, error)
	GetPrice(ctx context.Context, req *PriceReq) (any, error)
	GetTradeIdeas(ctx context.Context) (any, error)
}

type falconService struct {
	client       web.Client
	baseURL      string
	midasBaseURl string
}

// NewFalconService creates a new instance of FalconService
func NewFalconService(client web.Client) FalconService {
	return &falconService{
		client:       client,
		baseURL:      falconBaseURL,
		midasBaseURl: midasBaseURL,
	}
}

// PlaceOrder places a new order
func (s *falconService) PlaceOrder(ctx context.Context, req FalconRequest) (*Order, error) {
	order := &Order{
		UserOrder: UserOrder{
			OrderSource: 5,
			ModifyOrder: ModifyOrder{
				OrderItem: OrderItem{
					ExchangeName:    req.ExchangeName,
					Token:           req.Token,
					TradingSymbol:   req.TradingSymbol,
					Quantity:        req.Quantity,
					Price:           req.Price,
					TriggerPrice:    req.TriggerPrice,
					OrderType:       req.OrderType,
					TransactionType: req.TransactionType,
					PriceType:       req.PriceType,
					Validity:        req.Validity,
					DiscQuantity:    req.DiscQuantity,
					IsAMO:           req.IsAMO,
				},
				ProtectioParam: ProtectioParam{
					TargetPrice:   req.TargetPrice,
					StopLossPrice: req.StopLossPrice,
					TrailPrice:    req.TrailPrice,
				},
			},
		},
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize order: %w", err)
	}

	url := fmt.Sprintf("%s/v0/order/create/", s.baseURL)
	httpReq := web.NewHTTPRequest(url, http.MethodPost, orderJSON, web.WithHeaders("Authorization", internal.AuthToken))

	var resp Order
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to place order: %w", err)
	}
	return &resp, nil
}

// GetHoldings retrieves holdings for an account
func (s *falconService) GetHoldings(ctx context.Context, accountID string) (any, error) {
	url := fmt.Sprintf("%s/v1/report/holdings/", s.baseURL)
	httpReq := web.NewHTTPRequest(url, http.MethodGet, nil, web.WithHeaders("Authorization", internal.AuthToken))
	var resp any
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to get holdings: %w", err)
	}
	return resp, nil
}

// GetPositions retrieves positions for an account
func (s *falconService) GetPositions(ctx context.Context, accountID string) (any, error) {
	url := fmt.Sprintf("%s/v0/report/positions/", s.baseURL)
	httpReq := web.NewHTTPRequest(url, http.MethodGet, nil, web.WithHeaders("Authorization", internal.AuthToken))

	var resp any
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to get positions: %w", err)
	}
	return resp, nil
}

// GetSecurityInfo retrieves information about a security
func (s *falconService) GetSecurityInfo(ctx context.Context, req *SecurityInfoReq) (any, error) {
	url := fmt.Sprintf("%s/v0/security/%s", s.baseURL, req.Token)
	req.ExchangeName = 1
	jsonReq, _ := json.Marshal(req)
	httpReq := web.NewHTTPRequest(url, http.MethodPost, jsonReq, web.WithHeaders("Authorization", internal.AuthToken))

	var resp any
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to get security info: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetOrderBook(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v0/report/orders/", s.baseURL)
	httpReq := web.NewHTTPRequest(url, http.MethodGet, nil, web.WithHeaders("Authorization", internal.AuthToken))

	var resp any
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to get order book: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetPrice(ctx context.Context, req *PriceReq) (any, error) {
	req.Mode = 3
	url := fmt.Sprintf("%s/v1/stocks/quotes", s.baseURL)
	jsonReq, _ := json.Marshal(req)
	httpReq := web.NewHTTPRequest(url, http.MethodPost, jsonReq, web.WithHeaders("Authorization", internal.AuthToken))

	var resp any
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to get price: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetTradeIdeas(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v0/ideas/?status=2", s.midasBaseURl)
	httpReq := web.NewHTTPRequest(url, http.MethodGet, nil, web.WithHeaders("Authorization", internal.AuthToken))

	var resp any
	if err := web.CallRestAPIV1(ctx, httpReq, s.client, web.WithDestination(&resp), web.WithErrorFunc(web.CommonErrCheck)); err != nil {
		return nil, fmt.Errorf("failed to get trade ideas: %w", err)
	}
	return resp, nil
}
