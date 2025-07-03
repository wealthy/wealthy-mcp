// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package websocket

import (
	"context"
	"fmt"
	sync "sync"

	"github.com/gorilla/websocket"
)

var wealthyWebsocket *websocket.Conn = nil

var priceStore sync.Map

// Add global context and cancel function
var (
	msgCtx    context.Context
	msgCancel context.CancelFunc
)

func Connect(ctx context.Context, url string) error {
	// Cancel previous message processing goroutine if it exists

	// Check if existing connection is still alive
	if wealthyWebsocket != nil {
		if err := wealthyWebsocket.WriteMessage(websocket.PingMessage, nil); err != nil {
			// Connection is dead, set to nil so we can reconnect
			wealthyWebsocket = nil
			if msgCancel != nil {
				msgCancel()
				msgCancel = nil
			}
		} else {
			return nil
		}
	}
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %w", err)
	}
	wealthyWebsocket = conn

	// Create a new context for the new goroutine
	msgCtx, msgCancel = context.WithCancel(ctx)
	go processMessages(msgCtx)
	return nil
}

func SubscribePrice(ctx context.Context, token string) (any, error) {
	msg := &PriceSubscriptionReq{
		Operation: 1,
		Mode:      1,
		Symbol:    []string{token},
	}
	if wealthyWebsocket == nil {
		return nil, fmt.Errorf("websocket connection not established")
	}

	if err := wealthyWebsocket.WriteJSON(msg); err != nil {
		return nil, fmt.Errorf("failed to write to websocket: %w", err)
	}
	return nil, nil
}
func processMessages(ctx context.Context) {
	messages, err := readMessages(ctx)
	if err != nil {
		return
	}
	for msg := range messages {
		switch msg.Data.(type) {
		case *Message_Feed:
			fmt.Println(msg.Data)
		}
	}
}

func readMessages(ctx context.Context) (<-chan *Message, error) {
	if wealthyWebsocket == nil {
		return nil, fmt.Errorf("websocket connection not established")
	}

	messages := make(chan *Message)

	go func() {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg := &Message{}
				err := wealthyWebsocket.ReadJSON(msg)
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						// Log unexpected close errors
						continue
					}
					// Connection is closed, exit goroutine
					return
				}
				messages <- msg
			}
		}
	}()
	return messages, nil
}

func GetLTP(ctx context.Context, token string) (any, error) {
	price, ok := priceStore.Load(token)
	if !ok {
		return nil, fmt.Errorf("price not found")
	}
	return price, nil
}

func storePrice(ctx context.Context, token string, price any) {
	priceStore.Store(token, price)
}
