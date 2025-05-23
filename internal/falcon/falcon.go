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
	"sync"

	"github.com/wealthy/wealthy-mcp/internal"
)

var (
	falconBaseURL   = "https://api.wealthy.in/broking/api"
	midasBaseURL    = "https://api.wealthy.in/midas/api"
	searchURL       = "http://scout.wealthy.in/api/v0/search/?q=%s&pt=stocks"
	ErrUnauthorized = errors.New("unauthorized")
	wsURL           = "https://api.wealthy.in/broking/api/v0/auth/oms/token/"
	addToWlSuccess  = "Successfully udpated to watchlist"
)

// FalconRequest represents the common parameters for Falcon API requests

// FalconService defines the interface for Falcon API operations
type FalconService interface {
	//order
	PlaceOrder(ctx context.Context, req []OrderReq) ([]placeOrderResponse, error)
	//reports
	GetHoldings(ctx context.Context) (any, error)
	GetPositions(ctx context.Context) (any, error)
	GetOrderBook(ctx context.Context) (any, error)
	GetPrice(ctx context.Context, req *PriceReq) (any, error)
	//research
	GetTradeIdeas(ctx context.Context) (any, error)
	GetSecurityInfo(ctx context.Context, req *SecurityInfoReq) (any, error)
	//watchlist
	AddToWatchlist(ctx context.Context, req *WatchlistReq) (any, error)
	GetWatchlists(ctx context.Context) (any, error)
	CreateWatchlist(ctx context.Context, name string) (any, error)
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
func (s *falconService) PlaceOrder(ctx context.Context, req []OrderReq) ([]placeOrderResponse, error) {
	for i := range req {
		req[i].OrderSource = 5
	}
	orderJSON, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize order: %w", err)
	}

	url := fmt.Sprintf("%s/v0/order/basket/", s.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(orderJSON))
	if err != nil {
		slog.Error("failed to create request", "error", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp []placeOrderResponse

	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to place order: %w", err)
	}
	return resp, nil
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
	url := fmt.Sprintf("%s/v1/stock/quotes/", s.baseURL)
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
	url := fmt.Sprintf("%s/v0/idea/?status=2", s.midasBaseURl)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get trade ideas: %w, browse internet for trade ideas", err)
	}
	return resp, nil
}

func (s *falconService) GetSecurityInfo(ctx context.Context, req *SecurityInfoReq) (any, error) {
	url := fmt.Sprintf(searchURL, req.Name)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to search security: %w", err)
	}
	return resp, nil
}

func (s *falconService) GetWebsocketURL(ctx context.Context) (string, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, wsURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp WebsocketURLResponse
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return "", fmt.Errorf("failed to get websocket URL: %w", err)
	}
	return resp.BaseURL, nil
}

func (s *falconService) AddToWatchlist(ctx context.Context, req *WatchlistReq) (any, error) {
	url := fmt.Sprintf("%s/v0/watchlist/script/", s.baseURL)
	jsonReq, _ := json.Marshal(req)
	fmt.Println("add to watchlist err", string(jsonReq))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to add to watchlist: %w", err)
	}
	resp = addToWlSuccess
	return resp, nil
}

func (s *falconService) GetWatchlists(ctx context.Context) (any, error) {
	watchlistNames, err := s.getWatchlistNames(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get watchlist names: %w", err)
	}

	var userWatchlists []any
	var wg sync.WaitGroup
	for _, name := range watchlistNames {
		var mu sync.Mutex
		wg.Add(1)
		go func(n string) {
			var resp any
			defer wg.Done()
			url := fmt.Sprintf("%s/v0/watchlist/", s.baseURL)
			jsonReq, _ := json.Marshal(WatchlistReq{Name: n})
			httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonReq))
			if err != nil {
				return
			}
			httpReq.Header.Set("Authorization", internal.AuthToken)
			if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
				mu.Lock()
				userWatchlists = append(userWatchlists, map[string]any{n: nil})
				mu.Unlock()
				return
			}
			mu.Lock()
			userWatchlists = append(userWatchlists, map[string]any{n: resp})
			mu.Unlock()
		}(name)
	}
	wg.Wait()
	return userWatchlists, nil
}

func (s *falconService) getAllWatchlists(ctx context.Context) (any, error) {
	url := fmt.Sprintf("%s/v0/watchlist/", s.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get all watchlists: %w", err)
	}
	return resp, nil
}

func (s *falconService) getWatchlistNames(ctx context.Context) ([]string, error) {
	url := fmt.Sprintf("%s/v0/watchlist/", s.baseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp []string
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to get all watchlists: %w", err)
	}
	return resp, nil
}

func (s *falconService) CreateWatchlist(ctx context.Context, name string) (any, error) {
	url := fmt.Sprintf("%s/v0/watchlist/", s.baseURL)
	jsonReq, _ := json.Marshal(WatchlistReq{Name: name})
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Authorization", internal.AuthToken)

	var resp any
	if err := callRestAPI(ctx, httpReq, &resp, s.client); err != nil {
		return nil, fmt.Errorf("failed to create watchlist: %w", err)
	}
	return resp, nil
}
