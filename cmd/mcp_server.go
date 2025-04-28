// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/wealthy/mcp/internal"
	"github.com/wealthy/mcp/tools"
)

func newServer() *server.MCPServer {
	s := server.NewMCPServer(
		"wealthy-mcp",
		"0.1.0",
	)
	//add tools
	tools.AddFalconTool(s)
	prompt := mcp.NewPrompt("falcon-prompt",
		mcp.WithPromptDescription("wealthy mcp server prompts"))
	s.AddPrompt(prompt, server.PromptHandlerFunc(promptHandler))
	return s
}

func newGinServer() *gin.Engine {
	router := gin.Default()
	router.GET("/health", healthHandler)
	router.GET("/auth/callback/", internal.AuthHandler)
	return router
}

func run(transport, addr string, logLevel slog.Level) error {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))
	s := newServer()

	switch transport {
	case "stdio":
		router := newGinServer()
		srv := server.NewStdioServer(s)

		// Start HTTP server for auth in background
		go func() {
			if err := router.Run(addr); err != nil {
				slog.Error("HTTP server error", "error", err)
			}
		}()

		// Force user to login through browser
		if internal.AuthStage == internal.AUTH_NOT_STARTED || internal.AuthStage == internal.AUTH_FAILED {
			internal.BrowserLogin("http://" + addr + "/auth/callback")
		}

		slog.Info("Starting Wealthy MCP server using stdio transport")
		return srv.Listen(context.Background(), os.Stdin, os.Stdout)
	case "sse":
		router := newGinServer()
		srv := server.NewSSEServer(s,
			// server.WithSSEContextFunc(mcp.ExtractReqInformation),
			server.WithBasePath("/mcp"),
		)

		router.Any("/mcp/*path", gin.WrapF(srv.ServeHTTP))
		router.GET("/auth/callback/", internal.AuthHandler)
		router.GET("/health/", healthHandler)
		// Force user to login through browser
		var loginOnce sync.Once
		loginOnce.Do(func() {
			internal.BrowserLogin(addr + "/auth/callback")
		})
		slog.Info("Starting Wealthy MCP server using SSE transport", "address", addr)
		if err := router.Run(addr); err != nil {
			return fmt.Errorf("HTTP server error: %v", err)
		}

	default:
		return fmt.Errorf(
			"invalid transport type: %s. must be 'stdio' or 'sse'",
			transport,
		)
	}
	return nil
}

// healthHandler responds with a simple health status
func healthHandler(c *gin.Context) {
	c.String(200, "OK")
}

func main() {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.StringVar(
		&transport,
		"transport",
		"stdio",
		"Transport type (stdio or sse)",
	)
	addr := flag.String("sse-address", "localhost:8004", "The host and port to start the sse server on")
	logLevel := flag.String("log-level", "info", "Log level (debug, info, warn, error)")
	flag.Parse()

	if err := run(transport, *addr, parseLevel(*logLevel)); err != nil {
		panic(err)
	}
}

func parseLevel(level string) slog.Level {
	var l slog.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return slog.LevelInfo
	}
	return l
}
