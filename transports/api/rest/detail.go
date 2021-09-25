package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/transports/response"
	"github.com/spf13/cast"
)

func (t *handler) Detail(ctx *gin.Context) {
	var rsp response.Handler
	// sanitize input from request
	data, err := t.sanitizeRequestDetail(ctx)
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent to usecase after sanitize input
	result, err := t.usecase.Detail(ctx, data.(int))
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent output data to client
	rsp.Success(ctx, result)
}

func (t *handler) sanitizeRequestDetail(ctx *gin.Context) (interface{}, error) {
	RowID := cast.ToInt(ctx.Param("id"))
	return RowID, nil
}
