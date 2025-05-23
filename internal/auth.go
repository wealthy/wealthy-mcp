// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
)

const (
	AUTH_NOT_STARTED = iota
	AUTH_STARTED
	AUTH_SUCCESS
	AUTH_FAILED
)

var (
	loginURL   = "https://api.wealthy.in/wealthyauth/dashboard/login/"
	AuthToken  string
	httpClient = &http.Client{
		Timeout: 30 * time.Second,
	}
	AuthStage   int
	CallbackURL string

	DebugMode bool // Set from main if debug flag is enabled
)

func AuthRequired() bool {
	return AuthStage == AUTH_NOT_STARTED || AuthStage == AUTH_FAILED
}

func BrowserLogin(cbURL string) {
	// if DebugMode {
	// 	if f, err := os.Open("auth_token.txt"); err == nil {
	// 		defer f.Close()
	// 		buf := make([]byte, 4096)
	// 		n, _ := f.Read(buf)
	// 		token := string(buf[:n])
	// 		if len(token) > 0 {
	// 			AuthToken = token
	// 			AuthStage = AUTH_SUCCESS
	// 			fmt.Println("Loaded auth token from file in debug mode, skipping browser login.")
	// 			return
	// 		}
	// 	}
	// }
	if CallbackURL == "" {
		CallbackURL = cbURL
	}
	loginURL := fmt.Sprintf(loginURL+"?redirect_url=%s", url.QueryEscape(CallbackURL))
	fmt.Println("opening browser", loginURL)
	AuthStage = AUTH_STARTED
	err := browser.OpenURL(loginURL)
	if err != nil {
		log.Fatal("authentication failed", err)
	}
}

func AuthHandler(c *gin.Context) {
	authcode := c.Query("authorization_token")
	if authcode != "" {
		token, err := generateAuthToken(c.Request.Context(), authcode)
		if err != nil {
			AuthStage = AUTH_FAILED
			log.Fatal("authentication failed", err)
		}
		AuthToken = token
		AuthStage = AUTH_SUCCESS

		// If debug mode, save token to file
		if DebugMode {
			f, err := os.Create("auth_token.txt")
			if err == nil {
				f.WriteString(AuthToken)
				f.Close()
			} else {
				log.Println("Failed to write auth token to file:", err)
			}
		}
		// Return success with HTML that will close the window
		c.Header("Content-Type", "text/html")
		c.Status(200)
		c.Writer.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Authentication Successful</title>
				<script>
			setTimeout(function() {
				window.close();
			}, 1000);
		</script>
			</head>
			<body>
				<h1>Authentication Successful</h1>
				<p>You can close this window now.</p>
			</body>
			</html>
			`))
	}
}

func generateAuthToken(ctx context.Context, authCode string) (string, error) {
	url := "https://api.wealthy.in/wealthyauth/dashboard/fetch-internal-token-details/"
	reqBody := `{"authorization_token":"` + authCode + `"}`
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respBody map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}
	return respBody["access_token"].(string), nil
}
