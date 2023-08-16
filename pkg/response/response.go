package response

import (
	"github.com/gin-gonic/gin"
)


func Respond(code int, payload any, ctx *gin.Context) {
    ctx.JSON(code, payload)
}

func RespondError(code int, err error, ctx *gin.Context) {
	ctx.JSON(code, map[string]string{"error": err.Error()})
}

func RespondErrorMessage(code int, msg string, ctx *gin.Context) {
	ctx.JSON(code, map[string]string{"error": msg})
}
