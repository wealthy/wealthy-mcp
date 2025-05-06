package utils

import (
	"net/http"
	"time"

	"github.com/wealthy/wealthy-mcp/internal/falcon"
)

var FalconService = falcon.NewFalconService(&http.Client{Timeout: 10 * time.Second})
