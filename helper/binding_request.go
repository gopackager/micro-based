package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindingRequest(payload interface{}, ctx *gin.Context) error {
	bind := binding.Default(ctx.Request.Method, ctx.ContentType())
	return ctx.ShouldBindWith(&payload, bind)
}
