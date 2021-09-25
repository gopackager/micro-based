package rest

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/helper"
	"github.com/gopackager/micro-based/models"
	"github.com/gopackager/micro-based/transports/response"
)

func (t *handler) Create(ctx *gin.Context) {
	var payload models.UsersCreateInput
	var rsp response.Handler

	// binding request to struct
	if err := helper.BindingRequest(&payload, ctx); err != nil {
		log.Print("eeee", err)
		rsp.Failed(ctx, err)
		return
	}

	// sanitize input from request
	data, err := t.sanitizeRequestCreate(payload)
	if err != nil {
		log.Print("aaaa", err)
		rsp.Failed(ctx, err)
		return
	}

	// sent to usecase after sanitize input
	result, err := t.usecase.Create(ctx, data)
	if err != nil {
		log.Print("ssss", err)
		rsp.Failed(ctx, err)
		return
	}

	// sent output data to client
	rsp.Success(ctx, result)
}

func (t *handler) sanitizeRequestCreate(data models.UsersCreateInput) (interface{}, error) {
	data.Fullname = strings.ToLower(data.Fullname)
	data.Email = strings.ToLower(data.Email)
	if data.NewPassword != data.ConfirmPassword {
		return nil, errors.New("password must be equal")
	}
	return data, nil
}
