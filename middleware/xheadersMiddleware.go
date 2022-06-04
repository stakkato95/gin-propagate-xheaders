package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stakkato95/service-engineering-go-lib/logger"
)

const XHeadersKey = "x-headers"

func XHeadersPropagation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := http.Header{}

		for key, val := range ctx.Request.Header {
			if strings.ToLower(key)[0] == 'x' {
				headers.Set(key, val[0])
			}
		}

		ctx.Set(XHeadersKey, headers)
		ctx.Next()
	}
}

func GetXHeaders(ctx *gin.Context) http.Header {
	headersRaw, ok := ctx.Get(XHeadersKey)
	if !ok {
		logger.Fatal("XHeadersPropagation is not used")
	}
	headers, ok := headersRaw.(http.Header)
	if !ok {
		logger.Fatal("XHeadersPropagation cast error")
	}
	return headers
}
