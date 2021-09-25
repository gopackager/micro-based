package rest

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/helper"
	"github.com/gopackager/micro-based/models"
	"github.com/gopackager/micro-based/transports/response"
	"github.com/spf13/cast"
)

func (t *handler) Update(ctx *gin.Context) {
	var payload models.UsersUpdateInput
	var rsp response.Handler

	// binding request to struct
	if err := helper.BindingRequest(&payload, ctx); err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sanitize input from request
	data, err := t.sanitizeRequestUpdate(ctx, payload)
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent to usecase after sanitize input
	result, err := t.usecase.Update(ctx, data)
	if err != nil {
		rsp.Failed(ctx, err)
		return
	}

	// sent output data to client
	rsp.Success(ctx, result)
}

func (t *handler) sanitizeRequestUpdate(ctx *gin.Context, data models.UsersUpdateInput) (interface{}, error) {
	data.Fullname = strings.ToLower(data.Fullname)
	data.Email = strings.ToLower(data.Email)
	data.RowID = cast.ToInt(ctx.Param("id"))
	return data, nil
}
