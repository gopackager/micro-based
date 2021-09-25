package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/transports/response"
	"github.com/spf13/cast"
)

func (t *handler) Delete(ctx *gin.Context) {
	var rsp response.Handler

	// sanitize input from request
	data, err := t.sanitizeRequestDelete(ctx)
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent to usecase after sanitize input
	result, err := t.usecase.Delete(ctx, data.(int))
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent output data to client
	rsp.Success(ctx, result)
}

func (t *handler) sanitizeRequestDelete(ctx *gin.Context) (interface{}, error) {
	RowID := cast.ToInt(ctx.Param("id"))
	return RowID, nil
}
