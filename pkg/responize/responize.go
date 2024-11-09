package responize

import (
	"github.com/gin-gonic/gin"
	"github.com/mxmrykov/asterix-auth/internal/model"
)

func R(ctx *gin.Context, data interface{}, code int, message string, error bool) {
	ctx.JSON(code, model.Response{
		Payload: data,
		Status:  code,
		Message: message,
		Error:   error,
	})
}