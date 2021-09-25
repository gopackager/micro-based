package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/transports/response"
)

func (t *handler) Read(ctx *gin.Context) {
	var rsp response.Handler

	// sent to usecase after sanitize input
	result, err := t.usecase.Read(ctx)
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent output data to client
	rsp.Success(ctx, result)
}
