package falcon

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/wealthy/wealthy-mcp/internal"
)

func callRestAPI(ctx context.Context, httpReq *http.Request, resp any, client *http.Client) error {
	httpReq.Header.Set("Authorization", internal.AuthToken)

	httpResp, err := client.Do(httpReq)
	if err != nil {
		slog.Error("failed to call wealthy api", "error", err)
		return fmt.Errorf("network error: %w", err)
	}
	if httpResp.StatusCode == http.StatusUnauthorized {
		internal.BrowserLogin(internal.CallbackURL)
	}
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return fmt.Errorf("response status code: %d", httpResp.StatusCode)
	}

	if httpResp.StatusCode == http.StatusNoContent || httpResp.StatusCode == http.StatusCreated {
		return nil
	}
	defer httpResp.Body.Close()
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return fmt.Errorf("failed to decode response: %w, status code: %d", err, httpResp.StatusCode)
	}
	return nil
}
