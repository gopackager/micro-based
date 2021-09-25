package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/helper"
)

const headerXRequestID = "X-Request-ID"

func RequestID(ctx *gin.Context) {
	id := ctx.GetHeader(headerXRequestID)
	if id == "" {
		id = helper.GenerateUUID()
	}
	ctx.Header(headerXRequestID, id)
	ctx.Next()
}
