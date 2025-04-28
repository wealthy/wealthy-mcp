// Copyright (c) 2024 Wealthy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utils

import "context"

type ContextKey string

const (
	AuthTokenKey ContextKey = "authToken"
)

func GetAuthToken(ctx context.Context) string {
	if token, ok := ctx.Value(AuthTokenKey).(string); ok {
		return token
	}
	return ""
}
