package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Data interface{} `json:"records"`
	Err  string      `json:"errors,omitempty"`
}

func (h Handler) Success(ctx *gin.Context, result interface{}) {
	h.Data = result
	ctx.AbortWithStatusJSON(http.StatusOK, h)
}

func (h Handler) Failed(ctx *gin.Context, e error) {
	h.Err = e.Error()
	ctx.AbortWithStatusJSON(http.StatusBadRequest, h)
}
