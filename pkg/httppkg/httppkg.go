package httppkg

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseAuthToken(ctx *gin.Context) ([]string, error) {
	h := ctx.Request.Header["Authorization"]
	if h == nil && len(h) == 0 {
		return nil, errors.New("auth token is missing")
	}
	tokenHeader := h[0]
	if tokenHeader == "" {
		// Token is missing
		return nil, errors.New("auth token is missing")
	}
	splits := strings.Split(tokenHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		return nil, errors.New("token format is invalid")
	}
	return splits, nil

}
