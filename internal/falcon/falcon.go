// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package falcon

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/wealthy/wealthy-mcp/internal"
)

var (
	falconBaseURL   = "https://api.wealthy.in/broking/api"
	midasBaseURL    = "https://api.wealthy.in/midas/api"
	searchURL       = "http://scout.wealthy.in/api/v0/search/?q=%s&pt=stocks"
	ErrUnauthorized = errors.New("unauthorized")
)

// FalconRequest represents the common parameters for Falcon API requests

// FalconService defines the interface for Falcon API operations
type FalconService interface {
	PlaceOrder(ctx context.Context, req FalconRequest) (*Order, error)
	GetHoldings(ctx context.Context) (any, error)
	GetPositions(ctx context.Context) (any, error)
	GetOrderBook(ctx context.Context) (any, error)
	GetPrice(ctx context.Context, req *PriceReq) (any, error)
	GetTradeIdeas(ctx context.Context) (any, error)
	GetSecurityInfo(ctx context.Context, req *SecurityInfoReq) (any, error)
}

type falconService struct {
	client       *http.Client
	baseURL      string
	midasBaseURl string
}

// NewFalconService creates a new instance of FalconService
func NewFalconService(client *http.Client) FalconService {
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
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(orderJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp Order
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to place order: %w", err)
	}
	return &resp, nil
}

// GetHoldings retrieves holdings for an account
func (s *falconService) GetHoldings(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v1/report/holdings/", s.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get holdings: %w", err)
	}
	return resp, nil
}

// GetPositions retrieves positions for an account
func (s *falconService) GetPositions(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v0/report/positions/", s.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get positions: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetOrderBook(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v0/report/orders/", s.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get order book: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetPrice(ctx context.Context, req *PriceReq) (any, error) {
	req.Mode = 3
	url := fmt.Sprintf("%s/v1/stocks/quotes", s.baseURL)
	jsonReq, _ := json.Marshal(req)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get price: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetTradeIdeas(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v0/ideas/?status=2", s.midasBaseURl)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get trade ideas: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetSecurityInfo(ctx context.Context, req *SecurityInfoReq) (any, error) {
	url := fmt.Sprintf(searchURL, req.Name)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var resp []any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to search security: %w", err)
	}
	if len(resp) == 0 {
		return nil, fmt.Errorf("no security found")
	}
	return resp[0], nil
}

func callRestAPI(ctx context.Context, httpReq *http.Request, resp any, client *http.Client) error {
	httpReq.Header.Set("Authorization", internal.AuthToken)

	httpResp, err := client.Do(httpReq)
	if err != nil {
		slog.Error("failed to get trade ideas", "error", err)
		return fmt.Errorf("network error: %w", err)
	}
	if httpResp.StatusCode == http.StatusUnauthorized {
		internal.BrowserLogin(internal.CallbackURL)
	}
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return fmt.Errorf("response status code: %d", httpResp.StatusCode)
	}

	defer httpResp.Body.Close()
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	return nil
}
