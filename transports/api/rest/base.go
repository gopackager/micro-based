package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/usecases"
)

type Transports interface {
	Create(ctx *gin.Context)
	Read(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type handler struct {
	usecase usecases.Usecase
}

func New(uc usecases.Usecase) Transports {
	return &handler{
		usecase: uc,
	}
}
