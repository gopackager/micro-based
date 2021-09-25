package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/repositories"
)

type Usecase interface {
	Version(ctx *gin.Context, data interface{}) error
	Create(ctx *gin.Context, data interface{}) (interface{}, error)
	Read(ctx *gin.Context) (interface{}, error)
	Detail(ctx *gin.Context, data int) (interface{}, error)
	Update(ctx *gin.Context, data interface{}) (interface{}, error)
	Delete(ctx *gin.Context, data int) (interface{}, error)
}

type usecases struct {
	repo repositories.Repository
}

func New(repo repositories.Repository) Usecase {
	return &usecases{repo}
}
