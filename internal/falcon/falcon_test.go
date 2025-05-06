package falcon

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wealthy/wealthy-mcp/internal"
)

type mockResponse struct {
	Data interface{} `json:"data"`
}

func setupTestServer(t *testing.T, handler http.HandlerFunc) (*falconService, *httptest.Server) {
	server := httptest.NewServer(handler)
	client := &http.Client{}
	service := &falconService{
		client:       client,
		baseURL:      server.URL,
		midasBaseURl: server.URL,
	}
	return service, server
}

func TestPlaceOrder(t *testing.T) {
	tests := []struct {
		name    string
		req     FalconRequest
		want    *Order
		wantErr bool
	}{
		{
			name: "successful order placement",
			req: FalconRequest{
				OrderReq: OrderReq{
					ExchangeName:    1,
					Token:           "123",
					TradingSymbol:   "AAPL",
					Quantity:        100,
					Price:           "150.00",
					TransactionType: 1,
				},
			},
			want: &Order{
				UserOrder: UserOrder{
					OrderSource: 5,
					ModifyOrder: ModifyOrder{
						OrderItem: OrderItem{
							ExchangeName:    1,
							Token:           "123",
							TradingSymbol:   "AAPL",
							Quantity:        100,
							Price:           "150.00",
							TransactionType: 1,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid order",
			req: FalconRequest{
				OrderReq: OrderReq{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/v0/order/create/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// The response needs to match the exact structure expected by the service
				responseOrder := Order{
					UserOrder: UserOrder{
						OrderSource: 5,
						ModifyOrder: ModifyOrder{
							OrderItem: OrderItem{
								ExchangeName:    1,
								Token:           "123",
								TradingSymbol:   "AAPL",
								Quantity:        100,
								Price:           "150.00",
								TransactionType: 1,
							},
						},
					},
				}
				json.NewEncoder(w).Encode(responseOrder)
			})
			defer server.Close()

			got, err := service.PlaceOrder(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want.UserOrder.OrderSource, got.UserOrder.OrderSource)
			assert.Equal(t, tt.want.UserOrder.ModifyOrder.OrderItem.ExchangeName, got.UserOrder.ModifyOrder.OrderItem.ExchangeName)
			assert.Equal(t, tt.want.UserOrder.ModifyOrder.OrderItem.Token, got.UserOrder.ModifyOrder.OrderItem.Token)
			assert.Equal(t, tt.want.UserOrder.ModifyOrder.OrderItem.TradingSymbol, got.UserOrder.ModifyOrder.OrderItem.TradingSymbol)
			assert.Equal(t, tt.want.UserOrder.ModifyOrder.OrderItem.Quantity, got.UserOrder.ModifyOrder.OrderItem.Quantity)
			assert.Equal(t, tt.want.UserOrder.ModifyOrder.OrderItem.Price, got.UserOrder.ModifyOrder.OrderItem.Price)
			assert.Equal(t, tt.want.UserOrder.ModifyOrder.OrderItem.TransactionType, got.UserOrder.ModifyOrder.OrderItem.TransactionType)
		})
	}
}

func TestGetHoldings(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful holdings retrieval",
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"holdings": []interface{}{
						map[string]interface{}{
							"symbol": "AAPL",
							"qty":    float64(100),
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/v1/report/holdings/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetHoldings(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetPositions(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful positions retrieval",
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"positions": []interface{}{
						map[string]interface{}{
							"symbol": "AAPL",
							"qty":    float64(100),
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/v0/report/positions/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetPositions(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetOrderBook(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful order book retrieval",
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"orders": []interface{}{
						map[string]interface{}{
							"symbol": "AAPL",
							"qty":    float64(100),
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/v0/report/orders/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetOrderBook(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetPrice(t *testing.T) {
	tests := []struct {
		name    string
		req     *PriceReq
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful price retrieval",
			req: &PriceReq{
				Symbols: []string{"AAPL"},
			},
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"AAPL": map[string]interface{}{
						"price": "150.00",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/v1/stock/quotes/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetPrice(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTradeIdeas(t *testing.T) {
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful trade ideas retrieval",
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"ideas": []interface{}{
						map[string]interface{}{
							"symbol": "AAPL",
							"action": "BUY",
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/v0/idea/", r.URL.Path)
				assert.Equal(t, "2", r.URL.Query().Get("status"))

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetTradeIdeas(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetSecurityInfo(t *testing.T) {
	tests := []struct {
		name    string
		req     *SecurityInfoReq
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful security info retrieval",
			req: &SecurityInfoReq{
				Name: "AAPL",
			},
			want: map[string]interface{}{
				"stocks": []interface{}{
					map[string]interface{}{
						"exchange_name":   float64(3),
						"instrument_type": "EQUITY",
						"isin_number":     "INE0C5901022",
						"lot_size":        float64(80000),
						"name":            "AA Plus Tradelink Limited",
						"tick_size":       "0",
						"token":           "543319",
						"trading_symbol":  "AAPLUSTRAD",
					},
					map[string]interface{}{
						"exchange_name":   float64(3),
						"instrument_type": "EQUITY",
						"isin_number":     "INE493N01012",
						"lot_size":        float64(1),
						"name":            "HARIA APPARELS LTD",
						"tick_size":       "0",
						"token":           "538081",
						"trading_symbol":  "HARIAAPL",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Contains(t, r.URL.String(), tt.req.Name)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetSecurityInfo(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAddToWatchlist(t *testing.T) {
	tests := []struct {
		name    string
		req     *WatchlistReq
		want    interface{}
		wantErr bool
	}{
		{
			name: "successful watchlist addition",
			req: &WatchlistReq{
				Name: "My Watchlist",
				Scrips: []scrip{
					{
						Exchange: 1,
						Token:    "123",
					},
				},
			},
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"status": "success",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, "/v0/watchlist/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.AddToWatchlist(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetWatchlists(t *testing.T) {
	tests := []struct {
		name    string
		req     *WatchlistReq
		want    interface{}
		wantErr bool
	}{
		{
			name: "get specific watchlist",
			req: &WatchlistReq{
				Name: "My Watchlist",
			},
			want: map[string]interface{}{
				"data": map[string]interface{}{
					"name": "My Watchlist",
					"scrips": []interface{}{
						map[string]interface{}{
							"exchange": float64(1),
							"token":    "123",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "get all watchlists",
			req: &WatchlistReq{
				Name: "All",
			},
			want: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"name": "Watchlist 1",
					},
					map[string]interface{}{
						"name": "Watchlist 2",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				if tt.req.Name == "All" {
					assert.Equal(t, http.MethodGet, r.Method)
				} else {
					assert.Equal(t, http.MethodPost, r.Method)
				}
				assert.Equal(t, "/v0/watchlist/", r.URL.Path)

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			got, err := service.GetWatchlists(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetWebsocketURL(t *testing.T) {
	// Save the original auth token, stage, and wsURL, and restore them after the test
	origAuthToken := internal.AuthToken
	origAuthStage := internal.AuthStage
	origWsURL := wsURL
	defer func() {
		internal.AuthToken = origAuthToken
		internal.AuthStage = origAuthStage
		wsURL = origWsURL
	}()
	internal.AuthToken = "test-token"
	internal.AuthStage = internal.AUTH_SUCCESS

	tests := []struct {
		name    string
		want    WebsocketURLResponse
		wantErr bool
	}{
		{
			name: "successful websocket URL retrieval",
			want: WebsocketURLResponse{
				BaseURL: "wss://api.wealthy.in/ws",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/broking/api/v0/auth/oms/token/", r.URL.Path)
				assert.Equal(t, "test-token", r.Header.Get("Authorization"))

				if tt.wantErr {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(tt.want)
			})
			defer server.Close()

			// Update wsURL to use the test server's URL
			wsURL = server.URL + "/broking/api/v0/auth/oms/token/"

			got, err := service.GetWebsocketURL(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want.BaseURL, got)
		})
	}
}
