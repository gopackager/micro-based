package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/helper"
)

const headerXAppID = "X-App-ID"

var bodyXAppID = "go-frame-"

func AppID(ctx *gin.Context) {
	id := ctx.GetHeader(headerXAppID)
	if id == "" {
		id = bodyXAppID + helper.GenerateUUID()[0:8]
	}
	ctx.Header(headerXAppID, id)
	ctx.Next()
}
